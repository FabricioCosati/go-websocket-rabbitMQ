import { useEffect, useState } from "react";
import { MessageDto } from "../dtos/message";

export function useFetchMessages(lastJsonMessage: unknown): MessageDto[] {
    const [messages, setMessages] = useState<MessageDto[]>([])

    useEffect(() => {
        if (lastJsonMessage == null) {
            return
        }
        const message: MessageDto = JSON.parse(JSON.stringify(lastJsonMessage))
        setMessages([...messages, message])

    }, [lastJsonMessage])

    return messages
}