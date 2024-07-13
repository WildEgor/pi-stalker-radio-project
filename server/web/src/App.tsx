import { Outlet } from "react-router-dom";
import Player from "@/organisms/Player/Player.tsx";

const App = () => {
  return (
      <div className="h-screen flex flex-col ">
        <div className="flex flex-1 overflow-hidden relative">
          <main className="flex flex-col flex-1 p-3 overflow-y-auto">
            <Outlet />
          </main>
        </div>
        <Player />
      </div>
  );
};

export default App;
