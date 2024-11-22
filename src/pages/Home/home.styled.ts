import { Button } from "@mui/material";
import styled from "styled-components";

const size = {
    mobileS: '320px',
    mobileM: '375px',
    mobileL: '425px',
    maxSm: '540px',
    tablet: '768px',
    laptop: '1024px',
    xl: '1140px',
    laptopL: '1440px',
    desktop: '2560px'
}

export const device = {
    mobileS: `(min-width: ${size.mobileS})`,
    mobileM: `(min-width: ${size.mobileM})`,
    mobileL: `(min-width: ${size.mobileL})`,
    maxSm: `(min-width: ${size.maxSm})`,
    tablet: `(min-width: ${size.tablet})`,
    laptop: `(min-width: ${size.laptop})`,
    xl: `(min-width: ${size.xl})`,
    laptopL: `(min-width: ${size.laptopL})`,
    desktop: `(min-width: ${size.desktop})`,
    desktopL: `(min-width: ${size.desktop})`
};

export const HomepageContainer = styled.div`
    background-color: #212121;

    padding-left: 1.25rem;
    padding-right: 1.25rem;
    @media ${device.maxSm} {
        padding-left: 2.5rem;
        padding-right: 2.5rem;
    }
    @media ${device.xl} {
        padding-left: 10rem;
        padding-right: 10rem;
    }

    ::-webkit-scrollbar {
        width: 12px;
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

export const ChatContainer = styled.div`
   display: flex;
   flex-direction: column;
   justify-content: flex-end;
   height: 100vh;
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
    font-size: 16px;
    resize: none;
    overflow: hidden;
    box-sizing: border-box;
    &:focus{
        outline: none;
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