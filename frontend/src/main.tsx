import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router";
import FormsProvider from "./context/FormsProvider.tsx";
import App from "./App.tsx";
import Register from "./ui/pages/Register.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <FormsProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </BrowserRouter>
    </FormsProvider>
  </StrictMode>
);
