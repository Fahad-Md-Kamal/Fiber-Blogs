import { useForm } from "react-hook-form";
import {
  Button,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
  TextField,
} from "@mui/material";
import React, { useState } from "react";
import Box from "@mui/material/Box";
import FormControl from "@mui/material/FormControl";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";



export type SignupFormInputs = {
  email: string;
  username: string;
  password: string;
};


const Signup: React.FC = () => {
  // const navigate = useNavigate();
  const { register, handleSubmit } = useForm<SignupFormInputs>();
  const [showPassword, setShowPassword] = useState<boolean>(false);
  const handleClickShowPassword = () => setShowPassword((show) => !show);
  const handleMouseDownPassword = (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    event.preventDefault();
  };

  const onSubmit = async (data: SignupFormInputs) => {
    // if (await loginApiService(data)){
    //   navigate("/");
    // };
  };

  return (
    <Box
      onSubmit={handleSubmit(onSubmit)}
      className="flex justify-center items-center max-h-screen"
      component="form"
      sx={{ "& .MuiTextField-root": { m: 1, width: "25ch" } }}
      noValidate
      autoComplete="off"
    >
      <TextField
        id="username"
        label="Username"
        variant="outlined"
        type="text"
        {...register("username")}
      />
      <TextField
        id="username"
        label="Email"
        variant="outlined"
        type="text"
        {...register("email")}
      />
      <FormControl sx={{ m: 1, width: "25ch" }} variant="outlined">
        <InputLabel htmlFor="outlined-adornment-password">Password</InputLabel>
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
      <Button type="submit" variant="contained">
        Submit
      </Button>
    </Box>
  );
};

export default Signup;
