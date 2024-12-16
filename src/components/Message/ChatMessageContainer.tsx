import { MessageDto } from "../../dtos/message";
import { DropdownContainer, DropDownItem, DropDownMenu, MessageContainer, MessagesContainer, MessagesUserPhoto, MessageText, MessageUser } from "../../pages/Home/home.styled";
import 'bootstrap/dist/css/bootstrap.min.css';
import { useState } from "react";
import { UserDto } from "../../dtos/user";

type Props = {
    messages: MessageDto[]
    user: UserDto
}

export default function ChatMessageContainer({ messages, user }: Props) {
    const [messageIndex, setMessageIndex] = useState<number | null>(null);
    const [showUserOpt, setShowUserOpt] = useState<boolean>(false);

    function handleUserOpt(key: number) {
        setMessageIndex((prev) => (prev === key ? null : key))
        setShowUserOpt(true)
    }

    function getUserData(message: MessageDto) {
        setShowUserOpt(false)
    }

    return (
        <MessagesContainer>
            {messages && messages.map((message: MessageDto, index: number) => (
                <MessageContainer key={index} onClick={() => handleUserOpt(index)}>
                    {messageIndex === index && user.Id !== message.User.Id && (
                        <DropdownContainer show={showUserOpt} onToggle={(isOpen) => setShowUserOpt(isOpen)}>
                            <DropDownMenu show>
                                <DropDownItem onClick={() => getUserData(message)}>Enviar Mensagem</DropDownItem>
                            </DropDownMenu>
                        </DropdownContainer>
                    )}
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