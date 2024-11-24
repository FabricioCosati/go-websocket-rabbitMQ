import useWebSocket from "react-use-websocket"
import { ChatContainer, HomepageContainer } from "./home.styled"
import { useFetchMessages } from "../../hooks/useFetchMessages"
import ChatMessageContainer from "../../components/Message/ChatMessageContainer"
import ChatFormContainer from "../../components/Message/ChatFormContainer"

function Home() {
    const WS_URL = "ws://localhost:8080/ws"
    const { sendJsonMessage, lastJsonMessage } = useWebSocket(WS_URL)
    const messages = useFetchMessages(lastJsonMessage)

    return (
        <HomepageContainer>
            <ChatContainer>
                <ChatMessageContainer messages={messages} />
                <ChatFormContainer sendJsonMessage={sendJsonMessage} />
            </ChatContainer>
        </HomepageContainer>
    )
}

export default Home