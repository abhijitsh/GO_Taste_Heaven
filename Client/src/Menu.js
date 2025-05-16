import React, { useState } from "react";
import MenuCard from "./menuCard";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { Link } from "react-scroll";
// import{Link} from "react-router-dom"
import user from "./Images/user.png"
import "./Menu.css";
import data from "./data";
import Footer from "./Footer";

//Image


//SearchBar
import searchIcon from "./searchIcon.svg"
import Navbar from "./Navbar";

const Menu = ({updateUser }) => {

  const[query,setQuery]=useState("");
  const[resData, setResData]=useState(null)
  // const [filteredItems, setFilteredItems]=useState([]);
const userName = JSON.parse(localStorage.getItem("user")).name;
  const Logout = () => {
    localStorage.removeItem("user");
    const user = localStorage.getItem("user");
    toast.success("Thanks for Visiting Menu");
    updateUser(user);
  };
  // const [product,setProduct]=useState({
  //   id:null,
  //   title:"",
  //   description:"",
  //   price:"",
  // });
  const Search=(e) => {
    setQuery(e.target.value);
    console.log(query);
   
  
    }
    const filteredItems = data.filter((item)=>{
      return item.title.toLowerCase().includes(query.toLowerCase());
    })
    console.log(filteredItems)
  
  return (
    <>
     <Navbar inside={true} updateUser={updateUser}/>
      <p className="Menu-heading">
        Hey! <span className="Menu-user-name">{userName}</span>
        <br></br>Welcome to the Paradise, Take a deep dive into the heaven
      </p>
      <div style={{"backgroundColor":"#f6f5f9","padding":"2%","display":"flex","justifyContent":"center"}}>
        <input className="Menu-search-bar" type="search" placeholder="Search..." onChange={Search}/>
        <img className="Menu-searchIcon" src={searchIcon}/>
      </div>
      <div className="Menu-flexbox">
        {!filteredItems.length?<div className="Menu-404-notfound">404! Food Item Not Found</div>:filteredItems.map((food) => {
          return (
            <MenuCard
              title={food.title}
              url={food.url}
              desc={food.description}
              price={food.price}
              id={food.id}
            />
          );
        })}
      </div>
      <Footer />
    </>
  );
};

export default Menu;
