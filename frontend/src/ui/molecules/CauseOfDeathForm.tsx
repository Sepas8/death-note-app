import { useContext } from "react";
import { useForm, type SubmitHandler } from "react-hook-form";
import { FormsContext } from "../../context/FormsContext";
import type { FormsContextType } from "../../context/FormsProvider";
import Timer from "./Timer";

interface CauseOfDeathFormData {
  causeOfDeath: string;
}

const CauseOfDeathForm = () => {
  const { forms, setForms } = useContext(FormsContext) as {
    forms: FormsContextType;
    setForms: (forms: FormsContextType) => void;
  };

  const {
    register: registerCauseOfDeath,
    handleSubmit: handleSubmitCauseOfDeath,
  } = useForm<CauseOfDeathFormData>();

  const onSubmitCauseOfDeath: SubmitHandler<CauseOfDeathFormData> = (data) => {
    console.log("🚀 ~ App ~ data:", data);

    setTimeout(() => {
      setForms({
        ...forms,
        showCauseOfDeathForm: false,
        showDetailsForm: true,
      });
    }, 1000);
  };

  return (
    <div className="flex flex-col items-center justify-center gap-4 h-screen">
      <Timer
        initialMinutes={0}
        initialSeconds={40}
        onTimerEnd={handleSubmitCauseOfDeath(onSubmitCauseOfDeath)}
      />

      <form
        onSubmit={handleSubmitCauseOfDeath(onSubmitCauseOfDeath)}
        className="md:w-[450px] bg-dnGray p-4 flex flex-col gap-4 rounded-xl"
      >
        <label htmlFor="causeOfDeath">Causa de muerte</label>
        <textarea
          id="causeOfDeath"
          className="bg-dnBlack text-dnWhite p-2 rounded-md focus:outline-none"
          placeholder="Causa de muerte..."
          {...registerCauseOfDeath("causeOfDeath", {
            value: "Ataque al corazón",
          })}
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

export default CauseOfDeathForm;
