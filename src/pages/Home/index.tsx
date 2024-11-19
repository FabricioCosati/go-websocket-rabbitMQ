import { useEffect, useState } from "react"
import useWebSocket from "react-use-websocket"
import { MessageDto } from "../../dtos/message"
import { UserDto } from "../../dtos/user"

function Home() {
    const [messages, setMessages] = useState<MessageDto[]>([])
    const WS_URL = "ws://localhost:8080/ws"
    const { sendJsonMessage, lastJsonMessage } = useWebSocket(
        WS_URL,
    )

    function sendMessage(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()

        const message = (e.currentTarget.elements[0] as HTMLInputElement).value
        const user: UserDto = {
            name: "guest"
        }
        const messageToSend: MessageDto = {
            user: user,
            message: message
        }

        sendJsonMessage(messageToSend)
        e.currentTarget.reset()
    }

    useEffect(() => {
        if (lastJsonMessage == null) {
            return
        }
        const message: MessageDto = JSON.parse(JSON.stringify(lastJsonMessage))
        setMessages([...messages, message])

    }, [lastJsonMessage])

    return (
        <div className="homepage">
            <div className="chatDiv">
                {messages && messages.map((message: MessageDto, index: number) => (
                    <div key={index} className="messageDiv">
                        <span className="messageUser">{message.user.name} </span>
                        <span className="messageText">{message.message}</span>
                    </div>
                ))}
            </div>
            <form method="post" onSubmit={(e) => sendMessage(e)}>
                <label>
                    <input name="message" placeholder="your text here." />
                </label>
                <button type="submit">Submit Message</button>
            </form>
        </div>
    )
}

export default Home