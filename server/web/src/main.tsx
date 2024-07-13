import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import {
  Route,
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements,
} from "react-router-dom";
import RadioMap from "@/pages/RadioMap.tsx";
import {ThemeProvider} from "@/store/providers/theme-provider.tsx";
import {NavProvider} from "@/store/providers/navContext.tsx";
import {AudioProvider} from "@/store/providers/audioContext.tsx";
import {Toaster} from "@/atoms/toaster.tsx";

const router = createBrowserRouter(
    createRoutesFromElements(
        <Route path="/" element={<App />}>
          <Route path="/" element={<RadioMap />} />
        </Route>
    )
);

ReactDOM.createRoot(document.getElementById("root")!).render(
    <React.StrictMode>
      <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
        <NavProvider>
          <AudioProvider>
            <RouterProvider router={router} />
            <Toaster />
          </AudioProvider>
        </NavProvider>
      </ThemeProvider>
    </React.StrictMode>
);
