import { useEffect } from "react"
import useWebSocket from "react-use-websocket"

function Home() {
    const WS_URL = "ws://localhost:8080/ws"
    const { sendJsonMessage, lastJsonMessage } = useWebSocket(
        WS_URL,
    )

    useEffect(() => {
        if (lastJsonMessage == null) {
            return
        }
    }, [lastJsonMessage])

    function sendMessage(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()

        const message = (e.currentTarget.elements[0] as HTMLInputElement).value
        sendJsonMessage({ "message": message })
    }

    return (
        <div className="homepage">
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