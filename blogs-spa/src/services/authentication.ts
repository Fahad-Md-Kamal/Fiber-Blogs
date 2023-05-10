import { showFailureToaster, showSuccessToaster } from "../components/Tosters";
import { serializeResponseDataToUser } from "../types/authentication";

const PostApi = async <T>(urlPath: string, data: T, token?: string): Promise<T> => {
  try {
    const response = await fetch(`http://localhost:3000/${urlPath}`, {method: "POST",
      headers: { "Content-Type": "application/json", Authorization: `Bearer ${token}`},
      body: JSON.stringify(data),
    });
    let responseData = await response.json();
    if (!response.ok) {
      showFailureToaster(responseData.error);
      return responseData;
    }
    return responseData;
  } catch (error) {
    showFailureToaster("Failed to login");
  }
  return {} as T;
};

export const loginApiService = async <T>(data: T): Promise<boolean> => {

  const responseData = await PostApi("login/", data);
  const responseUser = serializeResponseDataToUser(responseData);

  const token = responseUser.token?.trim();

  if (token) {
    localStorage.setItem("blog-token", token);
    showSuccessToaster("Logged in successfully");
    return true;
  }

  return false;
};