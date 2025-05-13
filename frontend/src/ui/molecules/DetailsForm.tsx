import { useContext } from "react";
import { FormsContext } from "../../context/FormsContext";
import type { FormsContextType } from "../../context/FormsProvider";
import { useForm, type SubmitHandler } from "react-hook-form";
import Timer from "./Timer";

interface DeathDetailsFormData {
  deathDetails: string;
}

const DetailsForm = () => {
  const { forms, setForms } = useContext(FormsContext) as {
    forms: FormsContextType;
    setForms: (forms: FormsContextType) => void;
  };

  const {
    register: registerDeathDetails,
    handleSubmit: handleSubmitDeathDetails,
  } = useForm<DeathDetailsFormData>();

  const onSubmitDeathDetails: SubmitHandler<DeathDetailsFormData> = (data) => {
    console.log("🚀 ~ App ~ data:", data);

    setTimeout(() => {
      setForms({
        ...forms,
        showRegisterForm: true,
        showDetailsForm: false,
      });
    }, 1000);
  };

  return (
    <div className="flex flex-col items-center justify-center gap-4 h-screen">
      <Timer
        initialMinutes={6}
        initialSeconds={40}
        onTimerEnd={handleSubmitDeathDetails(onSubmitDeathDetails)}
      />

      <form
        onSubmit={handleSubmitDeathDetails(onSubmitDeathDetails)}
        className="md:w-[450px] bg-dnGray p-4 flex flex-col gap-4 rounded-xl"
      >
        <label htmlFor="deathDetails">Detalles de la muerte</label>
        <textarea
          id="deathDetails"
          className="bg-dnBlack text-dnWhite p-2 rounded-md focus:outline-none"
          placeholder="Detalles de la muerte..."
          {...registerDeathDetails("deathDetails")}
        />

        <button
          type="submit"
          className="bg-dnRed hover:bg-dnDarkRed text-dnWhite p-2 rounded-md cursor-pointer"
        >
          Registrar
        </button>
      </form>
    </div>
  );
};

export default DetailsForm;
