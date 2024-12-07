import { ChannelGroupContainer, ChannelImageContainer, ChannelNameContainer } from "../../pages/Home/home.styled";


export default function GroupContainer() {

    return (
        <ChannelGroupContainer>
            <ChannelImageContainer>
                <img src={require("../../imgs/usersGroup.png")} alt="guest-user" />
            </ChannelImageContainer>
            <ChannelNameContainer>
                <span>Chat Geral</span>
            </ChannelNameContainer>
        </ChannelGroupContainer>
    )
}