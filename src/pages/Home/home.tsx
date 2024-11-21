import { useEffect, useState } from "react"
import useWebSocket from "react-use-websocket"
import { MessageDto } from "../../dtos/message"
import { UserDto } from "../../dtos/user"
import { ChatContainer, HomepageContainer, MessagesContainer, SubmitChatButton, SubmitChatFormContainer, SubmitChatInput } from "./home.styled"

function Home() {
    const [text, setText] = useState<string>("")

    const [messages, setMessages] = useState<MessageDto[]>([])
    const WS_URL = "ws://localhost:8080/ws"
    const { sendJsonMessage, lastJsonMessage } = useWebSocket(
        WS_URL,
    )

    function sendMessage(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
        if (text === "") {
            return
        }

        const user: UserDto = {
            name: "guest"
        }
        const messageToSend: MessageDto = {
            user: user,
            message: text
        }

        sendJsonMessage(messageToSend)
        setText("")
    }

    useEffect(() => {
        if (lastJsonMessage == null) {
            return
        }
        const message: MessageDto = JSON.parse(JSON.stringify(lastJsonMessage))
        setMessages([...messages, message])

    }, [lastJsonMessage])

    function handleInput(e: React.KeyboardEvent<HTMLTextAreaElement>) {
        if (e.key === "Enter" && !e.shiftKey) {
            e.preventDefault()
            sendMessage(e as unknown as React.FormEvent<HTMLFormElement>)
        }

    }

    return (
        <HomepageContainer>
            <ChatContainer>
                <MessagesContainer>
                    {messages && messages.map((message: MessageDto, index: number) => (
                        <div key={index}>
                            <div className="messageUser">{message.user.name} </div>
                            <div className="messageText">{message.message}</div>
                        </div>
                    ))}
                </MessagesContainer>
                <form method="post" onSubmit={(e) => sendMessage(e)}>
                    <SubmitChatFormContainer>
                        <label>
                            <SubmitChatInput
                                name="message"
                                placeholder="seu texto aqui."
                                value={text}
                                onChange={(e) => setText(e.target.value)}
                                onKeyDown={handleInput} />
                        </label>
                        <SubmitChatButton type="submit" variant="contained">Enviar</SubmitChatButton>
                    </SubmitChatFormContainer>
                </form>
            </ChatContainer>
        </HomepageContainer>
    )
}

export default Home