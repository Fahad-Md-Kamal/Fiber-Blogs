import { ToastPosition, toast } from "react-toastify";


export const showSuccessToaster = (content: string, position?: ToastPosition | undefined) => {
    toast.success(content, {
        position: position ?? "bottom-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
        });
}

export const showFailureToaster = (content: string, position?: ToastPosition | undefined) => {
toast.error(content, {
    position: position ?? "bottom-right",
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
    progress: undefined,
    theme: "dark",
    });
}
