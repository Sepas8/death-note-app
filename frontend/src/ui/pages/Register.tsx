import { useCallback, useContext } from "react";
import { FormsContext } from "../../context/FormsContext";
import RegisterUserForm from "../molecules/RegisterUserForm";
import CauseOfDeathForm from "../molecules/CauseOfDeathForm";
import DetailsForm from "../molecules/DetailsForm";
import type { FormsContextType } from "../../context/FormsProvider";

const Register = () => {
  const { forms } = useContext(FormsContext) as { forms: FormsContextType };

  const renderForm = useCallback(() => {
    if (forms.showRegisterForm) return <RegisterUserForm />;
    if (forms.showCauseOfDeathForm) return <CauseOfDeathForm />;
    if (forms.showDetailsForm) return <DetailsForm />;
  }, [forms]);

  return (
    <div className="flex flex-col items-center justify-center gap-4 h-screen">
      {renderForm()}
    </div>
  );
};

export default Register;
