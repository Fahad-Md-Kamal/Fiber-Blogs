import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import {
  showFailureToaster,
  showSuccessToaster,
} from "../../components/Tosters";

type LoginFormInputs = {
  username: string;
  password: string;
};

function saveTokenToLocaStorage(token: string): boolean {
  localStorage.setItem("blog-token", token);
  localStorage.setItem("isLoggedin", "1");

  return true;
}

function Login() {
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<LoginFormInputs>();
  let responseData: any;

  const onSubmit = async (data: LoginFormInputs) => {
    try {
      const response = await fetch("http://localhost:3000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer <insert JWT token here>",
        },
        body: JSON.stringify(data),
      });
      responseData = await response.json();
      if (!response.ok) {
        showFailureToaster(responseData.error);
      } else if (saveTokenToLocaStorage(responseData.token)) {
        showSuccessToaster("Logged in successfully");
        navigate("/");
      }
    } catch (error) {
      console.log(error);
      showFailureToaster("Failed to login");
    }
  };

  return (
    <div className="flex justify-center items-center max-h-screen">
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
        <div className="flex flex-col">
          <label htmlFor="username">Username / Email</label>
          <input
            type="text"
            id="username"
            placeholder="fahadmdkamal@gmail.com"
            {...register("username")}
            className="border border-gray-300 rounded-md px-3 py-2"
          />
        </div>
        <div className="flex flex-col">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            placeholder="*****"
            id="password"
            {...register("password")}
            className="border border-gray-300 rounded-md px-3 py-2"
          />
        </div>
        <button
          type="submit"
          className="bg-blue-500 text-white rounded-md px-4 py-2 hover:bg-blue-600"
        >
          Log In
        </button>
      </form>
    </div>
  );
}

export default Login;
