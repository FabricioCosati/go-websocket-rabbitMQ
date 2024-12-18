import useWebSocket from "react-use-websocket"
import { ChannelsContainer, ChatContainer, HomeChatContainer, HomepageContainer } from "./home.styled"
import { useFetchMessages } from "../../hooks/useFetchMessages"
import ChatMessageContainer from "../../components/Message/ChatMessageContainer"
import ChatFormContainer from "../../components/Message/ChatFormContainer"
import { useEffect, useState } from "react"
import axios from "axios"
import { UserDto } from "../../dtos/user"
import GroupContainer from "../../components/Channel/GroupContainer"

function Home() {
    const WS_URL = "ws://localhost:8080/ws"
    const { sendJsonMessage, lastJsonMessage } = useWebSocket(WS_URL)
    const messages = useFetchMessages(lastJsonMessage)

    const [user, setUser] = useState<UserDto>({} as UserDto)
    useEffect(() => {
        getUser()
    }, [])
    useEffect(() => {
    }, [user])

    async function getUser() {
        const url = "http://localhost:8080/auth/guest"

        const alredyUser = window.sessionStorage.getItem("user")
        if (alredyUser) {
            const userDto: UserDto = JSON.parse(alredyUser)
            setUser(userDto)
            return
        }

        await axios.get(url).then((response) => {
            const userDto: UserDto = JSON.parse(response.data)
            setUser(userDto)
            window.sessionStorage.setItem("user", JSON.stringify(userDto))
        })
    }

    return (
        <HomepageContainer>
            <HomeChatContainer>
                <ChannelsContainer>
                    <GroupContainer />
                </ChannelsContainer>
                <ChatContainer>
                    <ChatMessageContainer messages={messages} user={user} />
                    <ChatFormContainer sendJsonMessage={sendJsonMessage} user={user} />
                </ChatContainer>
            </HomeChatContainer>
        </HomepageContainer>
    )
}

export default Home