import { useContext } from "react";
import { useForm, type SubmitHandler } from "react-hook-form";
import { FormsContext } from "../../context/FormsContext";
import type { FormsContextType } from "../../context/FormsProvider";

interface FormData {
  firstName: string;
  lastName: string;
  image: FileList;
}

const RegisterUserForm = () => {
  const { forms, setForms } = useContext(FormsContext) as {
    forms: FormsContextType;
    setForms: (forms: FormsContextType) => void;
  };

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>();

  const onSubmit: SubmitHandler<FormData> = (data) => {
    console.log("🚀 ~ App ~ data:", data);
    setTimeout(() => {
      setForms({
        ...forms,
        showRegisterForm: false,
        showCauseOfDeathForm: true,
      });
    }, 1000);
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="md:w-[450px] bg-dnGray p-4 flex flex-col gap-4 rounded-xl"
    >
      <label htmlFor="firstName">Nombre</label>
      <input
        type="text"
        id="firstName"
        className="bg-dnBlack text-dnWhite p-2 rounded-md focus:outline-none"
        {...register("firstName", { required: true })}
        placeholder="Nombre de la victima..."
        aria-invalid={!!errors.firstName}
      />
      {errors.firstName && (
        <p className="text-dnRed text-sm">*Este campo es requerido</p>
      )}
      <label htmlFor="lastName">Apellido</label>
      <input
        type="text"
        id="lastName"
        className="bg-dnBlack text-dnWhite p-2 rounded-md focus:outline-none"
        {...register("lastName", { required: true })}
        placeholder="Apellido de la victima..."
        aria-invalid={!!errors.lastName}
      />
      {errors.lastName && (
        <p className="text-dnRed text-sm">*Este campo es requerido</p>
      )}
      <label htmlFor="image">Imagen</label>
      <input
        type="file"
        multiple={false}
        className="bg-dnBlack text-dnWhite p-2 rounded-md focus:outline-none cursor-pointer"
        {...register("image", {
          required: true,
          validate: (value: FileList) => {
            if (!value) return "Este campo es requerido";

            const acceptedTypes = ["image/jpeg", "image/png"];
            const fileType = value[0].type;
            return !acceptedTypes.includes(fileType)
              ? "Solo se aceptan formatos JPG y PNG"
              : true;
          },
        })}
        accept="image/*"
        aria-invalid={!!errors.image}
        id="image"
      />
      {errors.image && (
        <p className="text-dnRed text-sm" id="image-error">
          {errors.image.message}
        </p>
      )}

      <button
        type="submit"
        className="bg-dnRed hover:bg-dnDarkRed text-dnWhite p-2 rounded-md cursor-pointer"
      >
        Registrar
      </button>
    </form>
  );
};

export default RegisterUserForm;
