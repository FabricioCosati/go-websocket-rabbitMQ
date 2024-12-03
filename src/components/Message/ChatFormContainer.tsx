import { useState } from "react";
import { SubmitChatButton, SubmitChatFormContainer, SubmitChatInput } from "../../pages/Home/home.styled";
import { UserDto } from "../../dtos/user";
import { MessageDto } from "../../dtos/message";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";

type Props = {
    sendJsonMessage: SendJsonMessage
    user: UserDto

}

export default function ChatFormContainer({ sendJsonMessage, user }: Props) {
    const [text, setText] = useState<string>("")

    function handleInput(e: React.KeyboardEvent<HTMLTextAreaElement>) {
        if (e.key === "Enter" && !e.shiftKey) {
            e.preventDefault()
            sendMessage(e as unknown as React.FormEvent<HTMLFormElement>)
        }
    }

    function sendMessage(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
        if (text === "") {
            return
        }

        const messageToSend: MessageDto = {
            User: user,
            Message: text
        }

        sendJsonMessage(messageToSend)
        setText("")
    }

    return (
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
    )
}