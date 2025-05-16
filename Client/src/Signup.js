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
import validator from "validator";

function Signup() {
  const navigate = useNavigate();
  const [data, setData] = useState({
    name: "",
    email: "",
    password: "",
    phone: 0,
  });
  const [validate, setValidate] = useState(true);
  let name, value;
  const handleData = (e) => {
    name = e.target.name;
    value = e.target.value;
    setData({ ...data, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log(data);
    if (!data.name || !data.email || !data.password || !data?.phone) {
      toast.warning("Please Fill the Data");
    } else {
      try {
        const res = await fetch("http://localhost:8000/api/signup", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });

        const newres = await res.json();
        console.log(newres);
        if (newres.status == 200 && newres.message == "User Already Exists") {
          toast.error("Id already exist, Please Login");

          // alert("Id already exist, Please Login");
          navigate("/login");
        } else {
          toast.success("Signed Up successfully");
          navigate("/login");
        }
      } catch (err) {
        console.log(err);
      }
    }
  };

  // console.log(dataname);
  return (
    <div>
      <div className="authpage-flexbox">
        <div className="authpage-left">
          <img src={leftImg} className="authpage-left-image" />
        </div>
        <div className="authpage-right">
          <Card className="authpage-right-div">
            <h1 className="auth-header">SignUp</h1>
            <div className="auth-input-div">
              <label htmlFor="name" name="name" className="auth-label">
                Name
              </label>
              <InputText
                id="name"
                onChange={handleData}
                name="name"
                aria-describedby="username-help"
                style={{ width: "100%" }}
                value={data?.name}
                required
              />
            </div>
            <div className="auth-input-div">
              <label htmlFor="username" className="auth-label">
                Email
              </label>
              <InputText
                id="username"
                name="email"
                onChange={handleData}
                aria-describedby="username-help"
                style={{ width: "100%" }}
                value={data?.email}
                required
              />
            </div>

            <div className="auth-input-div">
              <label htmlFor="password" className="auth-label">
                Password
              </label>
              <InputText
                id="password"
                onChange={handleData}
                name="password"
                aria-describedby="username-help"
                style={{ width: "100%" }}
                value={data?.password}
                required
              />
            </div>

            <div className="auth-input-div">
              <label htmlFor="phone" className="auth-label">
                Phone
              </label>
              <div className="p-inputgroup flex-1">
                <span className="p-inputgroup-addon">+91</span>
                <InputNumber
                  name="phone"
                  required
                  placeholder="Enter 10 digit Mobile Number"
                  useGrouping={false}
                  style={{ width: "100%" }}
                  value={data?.phone || ""}
                  onChange={(e) => {
                    setData((prev) => ({
                      ...prev,
                      phone: e.value,
                    }));
                    const val = validator.isMobilePhone(e.value + "", [
                      "en-IN",
                    ]);
                    setValidate(val);
                  }}
                />
                
              </div>
              {!validate && (
                  <Message
                    severity="error"
                    text="Please Enter a valid Phone-Number"
                    className="primereact-class"
                  />
                )}
            </div>

            <Button
              label="SignUp"
              onClick={handleSubmit}
              severity="success"
              className="auth-btn"
            ></Button>
            <p className="formfoot">
              Already Have an Account? <a href="/login">Please Login</a>
            </p>
          </Card>
        </div>
      </div>
    </div>
  );
}

export default Signup;
