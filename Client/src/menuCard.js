import React, { useState, useRef, useEffect, Component } from "react";
import "./menuCard.css";
import Img from "./Images/Testimonials/Michael.jpg";
import url from "./Images/Menu/Pakode.jpg";
import data from "./data";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function MenuCard(dish) {
  const [count, setCount] = useState(0);
  const [food, setFood] = useState({});
  const Inputelement = useRef(null);
  
  var arr = [];
  const Addtocart = async (dish) => {
    console.log(dish);
    const email=JSON.parse(localStorage.getItem('user')).email;
    const res = await fetch(`http://localhost:8000/api/cart?email=${email}&action=add`,{
      method: 'POST',
      headers: { 'Content-Type': 'application/json'},
      body: JSON.stringify({
        foodId : dish.id,
        quantity : 1,
        price : Number(dish.price)
      }),
    })
    const resnew= await res.json();
    console.log(resnew);
    setCount(count=>count+1);
    if(resnew.status == 200 && resnew.message=='Cart Updated successfully' || resnew.message == "Cart created successfully"){
      toast.success(`${dish.title} add to your cart`);
      //setResData(resnew.data);
    }
    else {
      toast.error(resnew.message);
    }
  };
  return (
    <div className="menu-card">
      <div className="menu-card-flexbox">
        <img className="menu-card-img" src={dish.url} />
        <div className="menu-card-content">
          <p className="menu-card-heading">{dish.title}</p>
          {dish.desc}
          <div className="menu-price-flexbox">
            <p className="menu-card-price"> ₹ {dish.price}/-</p>
            <div className="menu-card-btn-input">
              <input
                ref={Inputelement}
                className="menu-card-input"
                id={dish.id}
                placeholder={Inputelement ? count : "0"}
              />
              <button
                key={dish.id}
                className="menu-card-button"
                style={{ backgroundColor: "green" }}
                onClick={ () => Addtocart(dish)}
               
              >
                Add to Cart
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default MenuCard;
