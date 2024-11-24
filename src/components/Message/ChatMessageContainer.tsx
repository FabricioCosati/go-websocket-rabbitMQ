import { MessageDto } from "../../dtos/message";
import { MessageContainer, MessagesContainer, MessagesUserPhoto, MessageText, MessageUser } from "../../pages/Home/home.styled";

type Props = {
    messages: MessageDto[]
}

export default function ChatMessageContainer({ messages }: Props) {
    return (
        <MessagesContainer>
            {messages && messages.map((message: MessageDto, index: number) => (
                <MessageContainer key={index}>
                    <MessagesUserPhoto>
                        <img src={require("../../imgs/" + message.User.Photo)} alt="guest-user" />
                    </MessagesUserPhoto>
                    <div key={index}>
                        <MessageUser className="messageUser">{message.User.Name} <div>{message.Time}</div></MessageUser>
                        <MessageText className="messageText">{message.Message}</MessageText>
                    </div>
                </MessageContainer>
            ))}
        </MessagesContainer>
    )
}