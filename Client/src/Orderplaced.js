import React, { useEffect, useState } from "react";
import { Link } from "react-scroll";
import "./Orderplaced.css";
import Footer from "./Footer";
import Spinner from "./Spinner";
import { NavLink, useNavigate } from "react-router-dom";
import db from "./data";

//Toastify
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Navbar from "./Navbar";

function Orderplaced({updateUser}) {
  const [isLoading, setIsLoading] = useState(true);
  const [finalorder, setFinalOrder] = useState({});
  const { name, email, phone } = JSON.parse(localStorage.getItem("user"));
  const order_id = JSON.parse(localStorage.getItem("order_id"));
  console.log(order_id);
  const navigate = useNavigate();
  useEffect(() => {
    fetch(`http://localhost:8000/api/order?email=${email}&orderId=${order_id}`)
      .then((response) => response.json())
      .then((resNew) => {
        console.log("hemlo");
        console.log(resNew);
        setIsLoading(false);
        setFinalOrder(resNew?.data);
      });
  }, []);

  return (
    <>
      {isLoading ? (
        <Spinner title="Placing Order..." />
      ) : (
        <>
        <Navbar inside={true} updateUser={updateUser}/>
          <div className="orderplaced-page">
            <p className="orderplaced-header">
              Congratulations🎉 Your Order has been placed
            </p>
            <div className="orderplaced-box-1-parent">
              <div className="orderplaced-box-1">
                <p className="orderplaced-subheader-1">
                  Thank you for choosing us.
                </p>
                <p className="orderplaced-subheader-2">
                  Here is your Order Summary, keep getting the <b>Taste</b> of{" "}
                  <b>Heaven</b>
                </p>

                {finalorder && (
                  <>
                    <p className="orderplaced-subheader-3">Order Summary</p>
                    <div className="orderplaced-box-2">
                      <div className="orderplaced-flexbox">
                        <div className="orderplaced-order-box">
                          <p className="orderplaced-order-content">
                            Order Id: <b>{finalorder.orderid}</b>
                          </p>
                          <p className="orderplaced-order-content">
                            <b>Items:</b>
                            {finalorder.foods &&
                              finalorder.foods.map((singleFood) => {
                                return (
                                  <>
                                    <span>
                                      {" "}
                                      <b>{singleFood.quantity}</b> x{" "}
                                      {db[singleFood.foodId - 1]?.title},{" "}
                                    </span>
                                  </>
                                );
                              })}
                          </p>
                          <p className="orderplaced-order-content">
                            Amount: <b>Rs {finalorder.totalPrice}</b>
                          </p>
                          <p className="orderplaced-order-content">
                            Date: {finalorder?.date}
                          </p>
                          <p className="orderplaced-order-content">
                            Time: {finalorder?.time}
                          </p>
                        </div>
                        <div className="orderplaced-address-box">
                          <p className="orderplaced-address-content">{name}</p>
                          <p className="orderplaced-address-content">
                            {phone||"NA"}
                          </p>
                          <p className="orderplaced-address-content">
                            {finalorder?.address?.address}
                          </p>
                          <p className="orderplaced-address-content">
                            {finalorder?.address?.landmark}
                          </p>
                          <p className="orderplaced-address-content">
                            {finalorder?.address?.pin}
                          </p>
                        </div>
                      </div>
                    </div>
                  </>
                )}
              </div>
            </div>
          </div>
          <div style={{ backgroundColor: "#f5f6f9", padding: "2% 0%" }}>
            <p className="orderplaced-footer-content">
              Wanna Visit your Profile?
            </p>
              <button className="orderplaced-btn" onClick={()=> navigate('/user')}>
                Go
              </button>
          </div>
          <Footer className="orderplaced-footer" />
        </>
      )}
    </>
  );
}

export default Orderplaced;
