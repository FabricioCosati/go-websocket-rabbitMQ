import { Button } from "@mui/material";
import { Dropdown } from "react-bootstrap";
import styled from "styled-components";

const size = {
    mobileS: '320px',
    mobileM: '375px',
    mobileL: '425px',
    sm: '540px',
    tablet: '768px',
    laptop: '1024px',
    xl: '1140px',
    laptopL: '1440px',
    desktop: '2560px'
}

export const Mindevice = {
    mobileS: `(min-width: ${size.mobileS})`,
    mobileM: `(min-width: ${size.mobileM})`,
    mobileL: `(min-width: ${size.mobileL})`,
    sm: `(min-width: ${size.sm})`,
    tablet: `(min-width: ${size.tablet})`,
    laptop: `(min-width: ${size.laptop})`,
    xl: `(min-width: ${size.xl})`,
    laptopL: `(min-width: ${size.laptopL})`,
    desktop: `(min-width: ${size.desktop})`,
    desktopL: `(min-width: ${size.desktop})`
};

export const Maxdevice = {
    mobileS: `(max-width: ${size.mobileS})`,
    mobileM: `(max-width: ${size.mobileM})`,
    mobileL: `(max-width: ${size.mobileL})`,
    sm: `(max-width: ${size.sm})`,
    tablet: `(max-width: ${size.tablet})`,
    laptop: `(max-width: ${size.laptop})`,
    xl: `(max-width: ${size.xl})`,
    laptopL: `(max-width: ${size.laptopL})`,
    desktop: `(max-width: ${size.desktop})`,
    desktopL: `(max-width: ${size.desktop})`
};

export const HomepageContainer = styled.div`
    background-color: #212121;

    padding-left: 1.25rem;
    padding-right: 1.25rem;
    @media ${Mindevice.sm} {
        padding-left: 2.5rem;
        padding-right: 2.5rem;
    }
    @media ${Mindevice.xl} {
        padding-left: 10rem;
        padding-right: 10rem;
    }

    ::-webkit-scrollbar {
        width: 0.75rem;
    }

    ::-webkit-scrollbar-track {
        background: #2F2F2F;
        border-radius: .5rem;
    }

    ::-webkit-scrollbar-thumb {
        background: #1f1f1f;
        border-radius: .5rem;
    }

    ::-webkit-scrollbar-thumb:hover {
        background: #1b1b1b;
    }
`

// Separar em outro arquivo depois - tudo acima e abaixo

export const HomeChatContainer = styled.div`
    display: flex;
`

// Separar em outro arquivo depois - tudo acima e abaixo

export const ChannelsContainer = styled.div`
    width: 40%;
    background-color: #232323;
    
    @media ${Maxdevice.tablet} {
        display: none;
    }
`

export const ChannelGroupContainer = styled.div`
    display: flex;
    align-items: center;
    gap: 1rem;
    background-color: green;
    padding: .5rem;
    cursor: pointer;

    background-color: #1f1f1f;
`

export const ChannelImageContainer = styled.div`
    img {
        width: 30px;
        height: 30px;
        background-color: #ececec;
        border-radius: 50%;
        margin-top: 0.25rem;
    }
`

export const ChannelNameContainer = styled.div`
    span {
        color: white;
        font-weight: bold;
    }
`

// Separar em outro arquivo depois - tudo acima e abaixo

export const ChatContainer = styled.div`
   display: flex;
   flex-direction: column;
   justify-content: flex-end;
   height: 100vh;
   width: 100%;
   padding-left: 1rem;
`

// Separar em outro arquivo depois - tudo acima e abaixo

export const MessagesContainer = styled.div`
    width: 100%;
    max-height: 100%;
    color: white;
    word-wrap: break-word;
    word-break: break-word;
    overflow-wrap: break-word;
    white-space: normal;
    box-sizing: border-box;
    overflow-y: auto;
    font-size: 1rem;
`

export const MessageContainer = styled.div`
    display: flex;
    gap: 1rem;
    margin: 1rem;
    position: relative;
`

export const DropdownContainer = styled(Dropdown)`
    position: absolute;
    left: -5px;
`

export const DropDownMenu = styled(Dropdown.Menu)`
    background-color: #2f2f2f; 
`

export const DropDownItem = styled(Dropdown.Item)`
    color: white; 
    background-color: transparent; 

    &:hover {
        background-color: #2f2f2f; 
        color: #b8b8b8; 
    }

    &:active,  &.active,  &:focus {
        background-color: #2f2f2f;
        color: white;
        box-shadow: none;
    }
`

export const MessagesUserPhoto = styled.div`
    img {
        width: 40px;
        height: 40px;
        background-color: #ececec;
        border-radius: .5rem;
        margin-top: 0.25rem;
    }
`

export const MessageUser = styled.div`
    font-weight: bold;
    display: flex;
    gap: .5rem;

    div {
        font-weight: normal;
        font-size: 0.875rem;
        color: #c9c9c9;
    }
`

export const MessageText = styled.div`
`

// Separar em outro arquivo depois - tudo acima e abaixo

export const SubmitChatFormContainer = styled.div`  
    background-color : #2F2F2F;
    margin: 1rem 0 1rem 0;
    border-radius: 1rem;
    padding: .5rem;
    display: grid;
    grid-template-columns: 1fr auto;
    gap: 2rem;
    justify-content: space-between;
    align-items: center;
`
export const SubmitChatInput = styled.textarea`
    color: white;
    background: none;
    width: 100%;
    height: auto;
    padding-top: 8px;
    padding-bottom: 8px;
    border: 0px;
    border: none;
    font-size: 1rem;
    resize: none;
    overflow: hidden;
    box-sizing: border-box;
    &:focus{
        outline: none;
    }
    &::placeholder {
        color: #9c9c9c;
    }
`
export const SubmitChatButton = styled(Button)`
    && {
        background-color: #7600ad;

        &:hover {
            background-color: #8a00ca;
        }
    }
`