import React, { useState } from "react";
import { FormsContext } from "./FormsContext";
const response = await fetch(`${import.meta.env.VITE_API_URL}/api/people`);

export type FormsContextType = {
  showRegisterForm: boolean;
  showCauseOfDeathForm: boolean;
  showDetailsForm: boolean;
};

const FormsProvider = ({ children }: { children: React.ReactNode }) => {
  const [forms, setForms] = useState<FormsContextType>({
    showRegisterForm: true,
    showCauseOfDeathForm: false,
    showDetailsForm: false,
  });

  return (
    <FormsContext.Provider value={{ forms, setForms }}>
      {children}
    </FormsContext.Provider>
  );
};

export default FormsProvider;
