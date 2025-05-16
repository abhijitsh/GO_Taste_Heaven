import React, { useState, useEffect } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "primereact/resources/themes/lara-light-cyan/theme.css";
//Components
import Landingpage from "./Landingpage";
import Signup from "./Signup";
import Login from "./Login";
import Menu from "./Menu";
import Cart from "./cart";
import Confirmation from "./Confirmation";
import Orderplaced from "./Orderplaced";
import User from "./User";


function Routesnew() {
  const [user, setUser] = useState({});

  useEffect(() => {
    setUser(JSON.parse(localStorage.getItem("user")));
    console.log(user);
  }, []);

  const updateUser = (user) => {
    setUser(user);
    localStorage.setItem("user", JSON.stringify(user));
    // console.log(user);
  };

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Landingpage updateUser={updateUser} />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login updateUser={updateUser} />} />
        <Route
          path="/menu"
          element={
            user && user.email ? (
              <Menu updateUser={updateUser} />
            ) : (
              <Landingpage updateUser={updateUser} />
            )
          }
        />{" "}
        <Route
          path="/menu/cart"
          element={
            user && user.email ? (
              <Cart updateUser={updateUser} />
            ) : (
              <Landingpage updateUser={updateUser} />
            )
          }
        />
        <Route
          path="/menu/cart/confirm"
          element={
            user && user.email ? (
              <Confirmation updateUser={updateUser} />
            ) : (
              <Landingpage updateUser={updateUser} />
            )
          }
        />
        <Route
          path="/menu/cart/confirm/orderplaced"
          element={
            user && user.email ? (
              <Orderplaced updateUser={updateUser} />
            ) : (
              <Landingpage updateUser={updateUser} />
            )
          }
        />
        <Route
          path="/user"
          element={
            user && user.email ? (
              <User updateUser={updateUser} />
            ) : (
              <Landingpage updateUser={updateUser} />
            )
          }
        />
        <Route path="*" element={<Landingpage updateUser={updateUser} />} />
      </Routes>
    </BrowserRouter>
  );
}

export default Routesnew;
