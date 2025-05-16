import React, { useRef } from "react";
import "./Navbar.css";
import { Link } from "react-scroll";
import routes from "./routes";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { Avatar } from "primereact/avatar";
import user from "./Images/userImg.png";
import cart_icon from "./Images/Menu/icons8-cart-48.png";
import dish from './Images/dish.png';
function Navbar({ updateUser, inside }) {
  const userData = JSON.parse(localStorage.getItem("user"));
  return (
    <div className="navbar">
      <div className="left">
        <img src = {dish} style = {{height : '50px', width:'50px', margin:'5px 0px 5px 15px' }}/>
        <a href="/"> <div className="logo"> Taste Heaven</div></a> 
      </div>
      <ul className="right">
        <li>
          <a href="/"> Home</a>
        </li>
        {!inside && (
          <>
            {" "}
            <li>
              <Link
                to="about"
                spy={true}
                smooth={true}
                offset={-40}
                duration={500}
              >
                About
              </Link>
            </li>
            <li>
              <Link
                to="services"
                spy={true}
                smooth={true}
                offset={-40}
                duration={500}
              >
                Services
              </Link>
            </li>
            <li>
              <Link
                to="testimonial"
                spy={true}
                smooth={true}
                offset={-40}
                duration={500}
              >
                Testimonials
              </Link>
            </li>{" "}
          </>
        )}
        <li>
          <Link
            to="contact"
            spy={true}
            smooth={true}
            offset={-10}
            duration={500}
          >
            Contact
          </Link>
        </li>

        {userData?.email && (
          <li>
            <a href="/menu">Menu</a>
          </li>
        )}
        {inside && (
          <>
            <li>
              <a href="/menu/cart">
                <img className="menu-nav-cart-icon" src={cart_icon} />
              </a>
            </li>
            <li>
              <a href="/user">
                <Avatar
                  image={!userData?.url ? user : userData.url}
                  className="mr-2"
                  size="large"
                  shape="circle"
                />
                {/* <img
                  src={
                    userData?.picture == "" ? user : userData.picture
                  }
                /> */}
              </a>
            </li>
          </>
        )}
      </ul>
    </div>
  );
}

export default Navbar;
