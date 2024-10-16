import { useEffect, useState } from "react";
import useWebSocket from "react-use-websocket";
import { useStore } from "store";

import Modal from "components/modal/Modal";

const WS_URL = "ws://127.0.0.1:8000/api/v1/ws";

const READY_STATE_OPEN = 1;

const Table = () => {
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [selectedSourceIP, setSelectedSourceIP] = useState<string>("");
  const [selectedResponseBody, setSelectedResponseBody] = useState<string>("");

  const { shouldConnect, messages, addToMessages, updateIsConnected } =
    useStore();

  const handleModalOpen = (sourceIP: string, responseBody: string) => {
    setSelectedSourceIP(sourceIP);
    setSelectedResponseBody(responseBody);
    setIsModalOpen(true);
  };

  const handleModalClose = () => {
    setIsModalOpen(false);
  };

  const { lastMessage, readyState } = useWebSocket(
    WS_URL,
    {
      share: true,
      shouldReconnect: () => true,
      onOpen: () => {
        console.log("WebSocket connection opened!");
        updateIsConnected(true);
      },
      onClose: () => {
        console.log("WebSocket connection closed!");
        updateIsConnected(false);
      },
      onError: (event) => console.error("WebSocket error:", event),
      onMessage: (event) => console.log("Received message:", event.data),
    },
    shouldConnect // websocket switch
  );

  useEffect(() => {
    if (readyState === READY_STATE_OPEN && lastMessage) {
      addToMessages(lastMessage.data);
      console.log("Message -> " + lastMessage.data);
    }
  }, [lastMessage, readyState]);

  return (
    <div className="overflow-x-auto">
      <table className="table table-zebra">
        {/* head */}
        <thead>
          <tr>
            <th>Source</th>
            <th>Body</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {/* Map over messages to create a row for each message */}
          {messages.map((message, index) => {
            // Parse the message as JSON
            const parsedMessage = JSON.parse(message);

            // Extract the sourceIP and responseBody
            const source = parsedMessage.sourceIP || "Unknown";
            const body = parsedMessage.responseBody || "No body";

            return (
              <tr key={index}>
                <td>{source}</td> {/* Display the source IP */}
                <td>{body}</td> {/* Display the response body */}
                <td className="flex justify-normal">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className="size-6 cursor-pointer"
                    onClick={() => handleModalOpen(source, body)}
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z"
                    />
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"
                    />
                  </svg>

                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className="size-6 cursor-pointer ml-6"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M15.666 3.888A2.25 2.25 0 0 0 13.5 2.25h-3c-1.03 0-1.9.693-2.166 1.638m7.332 0c.055.194.084.4.084.612v0a.75.75 0 0 1-.75.75H9a.75.75 0 0 1-.75-.75v0c0-.212.03-.418.084-.612m7.332 0c.646.049 1.288.11 1.927.184 1.1.128 1.907 1.077 1.907 2.185V19.5a2.25 2.25 0 0 1-2.25 2.25H6.75A2.25 2.25 0 0 1 4.5 19.5V6.257c0-1.108.806-2.057 1.907-2.185a48.208 48.208 0 0 1 1.927-.184"
                    />
                  </svg>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>

      {/* Reusable Modal Component */}
      <Modal
        isOpen={isModalOpen}
        onClose={handleModalClose}
        sourceIP={selectedSourceIP}
        responseBody={selectedResponseBody}
      />
    </div>
  );
};

export default Table;
