import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import {
  Button,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
  TextField,
} from "@mui/material";
import { loginApiService } from "../../services/authentication";
import React, { useState } from "react";
import Box from "@mui/material/Box";
import FormControl from "@mui/material/FormControl";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";

export type LoginFormInputs = {
  username: string;
  password: string;
};

const Login: React.FC = () => {
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<LoginFormInputs>();
  const [showPassword, setShowPassword] = useState<boolean>(false);
  const handleClickShowPassword = () => setShowPassword((show) => !show);
  const handleMouseDownPassword = (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    event.preventDefault();
  };

  const onSubmit = async (data: LoginFormInputs) => {
    if (await loginApiService(data)) {
      navigate("/");
    }
  };

  return (
    <>
      <Box
        onSubmit={handleSubmit(onSubmit)}
        component="form"
        className="flex flex-col justify-center items-center max-h-screen"
        sx={{ "& .MuiTextField-root": { m: 1, width: "25ch" } }}
        noValidate
        autoComplete="off"
      >
        <TextField
          id="username"
          label="Email / Username"
          variant="outlined"
          type="text"
          {...register("username")}
        />
        <FormControl sx={{ m: 1, width: "25ch" }} variant="outlined">
          <InputLabel htmlFor="outlined-adornment-password">
            Password
          </InputLabel>
          <OutlinedInput
            id="outlined-adornment-password"
            type={showPassword ? "text" : "password"}
            {...register("password")}
            endAdornment={
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle password visibility"
                  onClick={handleClickShowPassword}
                  onMouseDown={handleMouseDownPassword}
                  edge="end"
                >
                  {showPassword ? <VisibilityOff /> : <Visibility />}
                </IconButton>
              </InputAdornment>
            }
            label="Password"
          />
        </FormControl>
        <p className="my-8">Don't have a account ? <Link to={`/auth/signup`}>Join</Link></p>
        <Button type="submit"  variant="contained">
          Submit
        </Button>
      </Box>
      </>
  );
};

export default Login;
