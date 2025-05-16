import React, { useState, useEffect } from "react";
import { NavLink } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import "./AuthPage.css";
import { Card } from "primereact/card";
import { InputText } from "primereact/inputtext";
import { InputNumber } from "primereact/inputnumber";
import { Button } from "primereact/button";
import { Message } from "primereact/message";
import leftImg from "./Images/AuthImg.png";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function SignIn({ updateUser }) {
  const navigate = useNavigate();
  const [userData, setUserData] = useState({ email: "", password: "" });
  let name, value;
  const loginData = (e) => {
    name = e.target.name;
    value = e.target.value;
    setUserData({ ...userData, [name]: value });
  };
  const Login = async (e) => {
    e.preventDefault();
    if (!userData.email || !userData.password) {
      toast.warning("Please Fill the Data");
    } else {
      const res = await fetch("http://localhost:8000/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      });
      const res1 = await res.json();
      console.log(res1)
      if (res1.status == 400 &&  res1.message == "User not found") {
        toast.error("Please SignUp first");
        navigate("/signup");
      } else if (res1.status == 400 && res1.message == "Invalid Password") {
        toast.error("Please Enter Correct Password");
        navigate("/login");
      } else if (res1.status == 200 && res1.message == "Login Successful") {
        console.log(res1.data)
        updateUser(res1.data);
        toast.success("LoggedIn Successfully");
        navigate("/menu");
      }
    }
  };
  return (
    <div>
      <div className="authpage-flexbox">
        <div className="authpage-left">
          <img src={leftImg} className="authpage-left-image" />
        </div>
        <div className="authpage-right">
          <Card className="authpage-right-div">
            <h1 className="auth-header">LogIn</h1>
            <div className="auth-input-div">
              <label htmlFor="username" className="auth-label">
                Email
              </label>
              <InputText
                id="username"
                name="email"
                onChange={loginData}
                aria-describedby="username-help"
                style={{ width: "100%" }}
                value={userData?.email}
                required
              />
            </div>

            <div className="auth-input-div">
              <label htmlFor="password" className="auth-label">
                Password
              </label>
              <InputText
                id="password"
                onChange={loginData}
                name="password"
                aria-describedby="username-help"
                style={{ width: "100%" }}
                value={userData?.password}
                required
              />
            </div>
            <Button
              label="Login"
              onClick={Login}
              severity="success"
              className="auth-btn"
            ></Button>
            <p className="formfoot">
              Don't Have an Account? <a href="/signup">Please SignUp</a>
            </p>
          </Card>
        </div>
      </div>
    </div>
  );
}

export default SignIn;
