import React, { useState, useEffect } from "react";
import { Link } from "react-scroll";
import "./Cart.css";
import Footer from "./Footer";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
//Images
import user from "./Images/user.png";
import cart from "./Images/Cart/Cart.svg";
import cart_icon from "./Images/Menu/icons8-cart-48.png";
import db from "./data";

//spinner
import Spinner from "./Spinner";
import Navbar from "./Navbar";
import { DataTable } from "primereact/datatable";
import { Column } from "primereact/column";
import { useNavigate } from "react-router-dom";
import { Button } from "primereact/button";
const datanew = [];
function Cart({ updateUser }) {
  const [isLoading, setIsLoading] = useState(true);
  const [totalPrice, setTotalPrice] = useState(0);
  const [data, setData] = useState([]);
  const [userAdress, setUserAddress] = useState();
  const { email, name } = JSON.parse(localStorage.getItem("user"));
  const navigate = useNavigate();
  useEffect(() => {
    fetch(`http://localhost:8000/api/cart?email=${email}`)
      .then((res) => {
        return res.json();
      })
      .then((resNew) => {
        console.log(resNew);
        setIsLoading(false);
        setData(resNew?.data?.foods);
        setTotalPrice(resNew?.data?.totalPrice);
      })
      .catch((err) => console.log(err));
      setUserAddress(JSON.parse(localStorage.getItem('address')));
  }, []);
  
  const imageBodyTemplate = (product) => {
    return (
      <img
        src={`${db[product.foodId - 1].url}`}
        alt={product.image}
        className="listrow-img"
      />
    );
  };

  const ActionBtn = async (action, id) => {
    try {
      const res = await fetch(
        `http://localhost:8000/api/cart?email=${email}&action=${action}`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            foodId: id,
            quantity: 1,
            price: Number(db[id - 1].price),
          }),
        }
      );
      const resnew = await res.json();
      console.log("add", resnew);
      if (
        resnew.status == 200 &&
        resnew.message == "Cart Updated successfully"
      ) {
        action == "add"
          ? toast.success(`1 more ${db[id - 1].title} added to your cart`)
          : toast.success(`1 ${db[id - 1].title} removed from your cart`);
        setData(resnew.data?.foods);
        setTotalPrice(resnew.data?.totalPrice);
      } else {
        toast.error(resnew.message);
      }
    } catch (err) {
      toast.error(err);
      throw new Error(err);
    }
  };

  // {data && data.map((singleFood)=>{
  //    setWholeTotal(wholeTotal+singleFood.price*singleFood.quantity);
  // })}
  console.log(totalPrice);
  console.log(data);
  if (!data || data.length == 0) {
    console.log(isLoading);

    return (
      <>
        {isLoading ? (
          <Spinner title="Preparing Your Cart..." />
        ) : (
          <>
            <Navbar inside={true} updateUser={updateUser} />
            <div
              style={{ backgroundColor: "#f5f6f9", height: "100vh" }}
              className="header-flexbox"
            >
              {
                <p className="cart-header">
                  Hey!{" "}
                  <span style={{ color: "blueviolet", fontWeight: "600" }}>
                    {name}
                  </span>
                  ,<br></br>Welcome to the Cart!,<br></br>
                  <span style={{ color: "red", fontWeight: "600" }}>
                    No Items yet
                  </span>{" "}
                </p>
              }

              <img
                style={{ backgroundColor: "#f5f6f9" }}
                className="cart-img"
                src={cart}
              />
            </div>
          </>
        )}
      </>
    );
  } else {
    return (
      <>
        {isLoading ? (
          <Spinner title="Preparing Your Cart..." />
        ) : (
          <>
            <Navbar inside={true} updateUser={updateUser} />
            <div className="header-flexbox">
              {name && (
                <p className="cart-header">
                  Hey!{" "}
                  <span style={{ color: "blueviolet", fontWeight: "600" }}>
                    {name}
                  </span>
                  ,<br></br> Welcome to the Cart!
                </p>
              )}

              <img className="cart-img" src={cart} />
            </div>
            <div className="card" style={{backgroundColor : "#f6f5f9", width : '100%'}}>
              <DataTable value={data} tableStyle={{ maxWidth: "90%" }}>
                <Column field="foodId" header="ID"></Column>
                {/* <Column field="item" header="Item" body = {(prod) => {return <img src = {db[prod?.foodId - 1].url} className="listrow-img"/>}}></Column> */}
                <Column header="Item" body={imageBodyTemplate}></Column>
                <Column
                  field="name"
                  header="Name"
                  body={(prod) => {
                    return db[prod?.foodId - 1].title;
                  }}
                ></Column>
                <Column
                  field="quantity"
                  header="Quantity"
                  body={(prod) => {
                    return (
                      <div className="cart-quantity-flexbox">
                        <button
                          className="cart-quantity-btn"
                          onClick={() => ActionBtn("remove", prod.foodId)}
                        >
                          -
                        </button>
                        <p className="cart-quantity">{prod.quantity}</p>
                        <button
                          className="cart-quantity-btn"
                          onClick={() => ActionBtn("add", prod.foodId)}
                        >
                          +
                        </button>
                      </div>
                    );
                  }}
                ></Column>
                <Column
                  field="price"
                  header="Price"
                  body={(prod) => {
                    return Number(db[prod?.foodId - 1].price).toLocaleString(
                      "en-US",
                      { style: "currency", currency: "INR" }
                    );
                  }}
                ></Column>

                <Column
                  field="total"
                  header="Total"
                  body={(prod) => {
                    return (
                      Number(db[prod?.foodId - 1].price) * prod.quantity
                    ).toLocaleString("en-US", {
                      style: "currency",
                      currency: "INR",
                    });
                  }}
                ></Column>

                {/* <Column field="category" header="Category"></Column>
                <Column field="quantity" header="Quantity"></Column> */}
              </DataTable>
              <div className="cart-totalprice-box">
                <p className="cart-totalprice">Total Amount : ₹ {totalPrice}</p>
              </div>
            </div>

            <div className="cart-totalprice-box-parent">
           
             {!userAdress?.address?<h2 style={{textAlign : 'center', margin:'6vh'}}>Please Add Address to Continue....</h2> : <div className="user-address-tag-card" style={{width : 'fit-content', margin:'2vh auto'}}>
                <p style={{textAlign:'center', fontSize :'1.5rem'}}><b>Address</b></p>
                <div className="user-address-card">
                  <div className="order-card-address">
                    <p className="address-card-address-text">
                      {userAdress?.address}
                    </p>
                    <p className="address-card-address-text">
                      {userAdress?.landmark}
                    </p>
                    <p className="address-card-address-text">
                      {userAdress?.pin}
                    </p>
                  </div>
                </div>
              </div>}
             
              <Button onClick={() =>{navigate('/menu/cart/confirm')}} className ="checkout-btn" type='submit' label="Checkout" severity="help" disabled = {!userAdress?.address}></Button>
            </div>
            <Footer />
          </>
        )}
      </>
    );
  }
}

export const food = datanew;
export default Cart;
