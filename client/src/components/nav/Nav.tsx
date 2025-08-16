import { Radio, Zap, ZapOff } from "lucide-react";
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
        <a className="link no-underline text-xl font-bold flex items-center justify-center cursor-pointer">
          <Radio size={32} />
          <span className="ml-2">Skiza</span>
        </a>
        <span className="mt-2">Postback Receiver and Stream</span>
      </div>
      <div className="shrink flex items-center justify-center">
        {isConnected && <Zap fill="green" strokeWidth={0} size={32} />}
        {!isConnected && <ZapOff fill="red" strokeWidth={0} size={32} />}
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
