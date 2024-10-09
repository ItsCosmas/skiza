import { useState } from "react";
import { useStore } from "store";

function Nav() {
  const { updateShouldConnect, isConnected } = useStore();

  const [switchHovered, setSwitchHovered] = useState(false);
  const [wasMouseOutside, setWasMouseOutside] = useState(true); // Tracks if the mouse was outside before entering

  const handleConnectionSwitchHover = () => {
    // Only toggle hover state if the connection is active and mouse was previously outside
    if (isConnected && wasMouseOutside) {
      setSwitchHovered(!switchHovered);
      setWasMouseOutside(false); // Set to false once the mouse enters
    }
  };

  const handleMouseLeave = () => {
    setWasMouseOutside(true); // Reset when the mouse leaves
    setSwitchHovered(false); // Reset hover state when mouse leaves
  };

  const handleConnectionSwitch = () => {
    updateShouldConnect(!isConnected);
    setWasMouseOutside(true); // Reset mouse state when connection changes
    setSwitchHovered(false); // Reset hover state when connection changes
    // Trigger Creation of WebSocket Connection
  };

  return (
    <nav className="navbar justify-between bg-base-100">
      <div className="shrink flex-col h-10 items-start prose">
        <a className="link no-underline text-xl font-bold flex justify-center cursor-pointer">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="size-6 mr-4"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M9.348 14.652a3.75 3.75 0 0 1 0-5.304m5.304 0a3.75 3.75 0 0 1 0 5.304m-7.425 2.121a6.75 6.75 0 0 1 0-9.546m9.546 0a6.75 6.75 0 0 1 0 9.546M5.106 18.894c-3.808-3.807-3.808-9.98 0-13.788m13.788 0c3.808 3.807 3.808 9.98 0 13.788M12 12h.008v.008H12V12Zm.375 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Z"
            />
          </svg>
          <span>Skiza</span>
        </a>
        <span className="mt-2">Postback Receiver and Stream</span>
      </div>
      <div className="shrink">
        {isConnected && (
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="green"
            className="size-6"
          >
            <path
              fillRule="evenodd"
              d="M14.615 1.595a.75.75 0 0 1 .359.852L12.982 9.75h7.268a.75.75 0 0 1 .548 1.262l-10.5 11.25a.75.75 0 0 1-1.272-.71l1.992-7.302H3.75a.75.75 0 0 1-.548-1.262l10.5-11.25a.75.75 0 0 1 .913-.143Z"
              clipRule="evenodd"
            />
          </svg>
        )}

        {!isConnected && (
          <>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="red"
              className="size-6"
            >
              <path d="m20.798 11.012-3.188 3.416L9.462 6.28l4.24-4.542a.75.75 0 0 1 1.272.71L12.982 9.75h7.268a.75.75 0 0 1 .548 1.262ZM3.202 12.988 6.39 9.572l8.148 8.148-4.24 4.542a.75.75 0 0 1-1.272-.71l1.992-7.302H3.75a.75.75 0 0 1-.548-1.262ZM3.53 2.47a.75.75 0 0 0-1.06 1.06l18 18a.75.75 0 1 0 1.06-1.06l-18-18Z" />
            </svg>
          </>
        )}
        <span className="ml-2">
          {isConnected ? "Connected" : "Disconnected"}
        </span>
        <input
          type="checkbox"
          // className="ml-4 toggle toggle-success"
          className={`ml-4 toggle ${
            switchHovered ? "toggle-error" : "toggle-success"
          }`}
          checked={isConnected}
          onChange={handleConnectionSwitch}
          onMouseEnter={handleConnectionSwitchHover}
          onMouseLeave={handleMouseLeave} // Handle mouse leave
        />
      </div>
    </nav>
  );
}

export default Nav;
