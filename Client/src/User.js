import React, { useState, useEffect } from "react";
import "./User.css";
import { Avatar } from "primereact/avatar";
import { Card } from "primereact/card";
import { TabView, TabPanel } from "primereact/tabview";
import "primereact/resources/themes/lara-light-indigo/theme.css";
import "primeicons/primeicons.css";
import { Tag } from "primereact/tag";
import { Button } from "primereact/button";
import { InputText } from "primereact/inputtext";
import { InputNumber } from "primereact/inputnumber";
import Navbar from "./Navbar";
import Spinner from "./Spinner";
import { Link } from "react-scroll";
import { useNavigate } from "react-router-dom";
import Footer from "./Footer";
import db from "./data";
import { Dialog } from "primereact/dialog";
import { InputTextarea } from "primereact/inputtextarea";
//Images
import edit from "./Images/edit.png";
import user from "./Images/userImg.png";
//Toastify
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const User = ({ updateUser }) => {
  const navigate = useNavigate();

  const [order, setOrder] = useState({});
  const userData = JSON.parse(localStorage.getItem("user"));
  const [nameuser, setNameUser] = useState(userData?.name);
  const [phone, setPhone] = useState(Number(userData?.phone));
  const emailuser = userData?.email;
  const [contactuser, setContactUser] = useState(userData?.phone || "NA");
  const [disableInput, setDisableInput] = useState(true);
  const [visible, setVisible] = useState(false);
  const [address, setAddress] = useState({ address: "", landmark: "", pin: 0 });
  const [addressRes, setAddressRes] = useState([]);
  const Logout = () => {
    localStorage.removeItem("user");
    const user = localStorage.getItem("user");
    localStorage.removeItem("address");
    toast.success("Thanks for Visiting Menu");
    updateUser(user);
  };
  const email = userData.email;
  const [profile, setProfile] = useState({
    name: nameuser,
    phone: userData?.phone,
  });
  useEffect(() => {
    fetch(`http://localhost:8000/api/profile?email=${email}`)
      .then((res) => res.json())
      .then((res1) => {
        console.log(res1);
        setLoading(false);
        setOrder(res1.data);
      });
  }, []);

  useEffect(() => {
    fetch(`http://localhost:8000/api/address?email=${email}`)
      .then((res) => res.json())
      .then((res1) => {
        console.log(res1);
        setLoading(false);
        setAddressRes(res1.data);
        const checkAddress = JSON.parse(localStorage.getItem("address"))
        if(!checkAddress && res1.data && res1.data[0]) handleHomeAddress(res1.data[0], 0);
      });
  }, []);
  // console.log(order);

  const [data, setData] = useState({});
  const [picture, setPicture] = useState(null);
  const [imagePreview, setImagePreview] = useState(null);
  const [loading, setLoading] = useState(true);
  const [editMode, seteditMode] = useState(false);
  const [homeTag, setHomeTag] = useState(0);

  const postDetails = async (picture) => {
    // console.log(picture);
    console.log(picture);
    setLoading(true);
    if (picture == undefined) {
      toast.warning("Please Upload a image");
      return;
    } else if (picture.size >= 1048576) {
      toast.warning("The size of image is greater than 1mb");
      setLoading(false);
      return;
    }
    const data = new FormData();
    data.append("file", picture);
    data.append("upload_preset", "tc3augsj");
    try {
      let res = await fetch(
        "https://api.cloudinary.com/v1_1/dcbrlaot1/image/upload",
        {
          method: "post",
          body: data,
        }
      );
      const urlData = await res.json();
      // console.log(urlData.url);
      setPicture(urlData.url.toString());
      setLoading(false);
      setImagePreview(URL.createObjectURL(picture));
    } catch (error) {
      console.log(error);
      setLoading(false);
    }
  };

  const handleImage = async (e) => {
    e.preventDefault();
    const res = await fetch(`http://localhost:8000/api/image?email=${email}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url: picture }),
    });
    const resNew = await res.json();
    if (
      resNew.status == 200 &&
      resNew.message == "Image Updated Successfully"
    ) {
      setLoading(false);
      toast.success("Image uploaded successfully");
      userData.url = picture;
      localStorage.setItem("user", JSON.stringify(userData));
      setImagePreview(null);
      setPicture(null);
    } else {
      setLoading(false);
      toast.error("Try again!");
    }
  };
  const showPopup = () => {
    toast.warn("Tap on Image to change !");
  };

  const handleEdit = () => {
    seteditMode(!editMode);
  };

  const handleNameChange = (e) => {
    setNameUser(e.target.value);
  };
  //  const handleEmailChange =  (e) => {
  // setContactUser(e.target.value);

  //  }
  const handleContactChange = (e) => {
    setContactUser(e.target.value);
  };
  const handleProfileSave = async () => {
    if (!profile?.name && !profile?.phone) {
      toast.error("Please make sure you have edited your profile");
      return;
    }
    console.log(profile);
    const res = await fetch(
      `http://localhost:8000/api/profile?email=${email}`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: profile.name,
          phone: Number(profile.phone),
        }),
      }
    );
    const res1 = await res.json();
    if (
      res1.status == 200 &&
      res1.message == "User Profile Updated Successfully"
    ) {
      setDisableInput(true);
      localStorage.setItem("user", JSON.stringify(res1.data));
      toast.success("Profile Updated");
      setNameUser(res1.data.name);
      setPhone(res1.data.phone);
    } else {
      setDisableInput(false);
      toast.error("Something went wrong", res1.message);
      setProfile({ name: "", phone: 0 });
    }
  };

  const handleAddress = (e) => {
    let naam = e?.target?.name;
    let val = e?.target?.value;

    setAddress((prev) => {
      return {
        ...prev,
        [naam]: val,
      };
    });
  };

  const handleAddressSave = async (e) => {
    e.preventDefault();
    console.log(address);
    if (!address?.address || !address?.landmark || !address?.pin) {
      toast.error("Please fill all the fields");
      return;
    }

    if (address?.pin.toString().length != 6) {
      toast.error("Please enter a valid pincode");
      return;
    }
    let res = await fetch(`http://localhost:8000/api/address?email=${email}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(address),
    });

    res = await res.json();

    if (res.status == 200 && res.message == "Address Saved Successfully") {
      toast.success("Address Save Successfully");
      console.log(addressRes, res.data)
      if(!addressRes || (addressRes &&  !addressRes?.length)) {
        setAddressRes([res.data]);
        handleHomeAddress(res.data, 0);
      }
      else setAddressRes((address => ([...address, res.data])))
      
      setAddress({ address: "", landmark: "", pin: 0 });
      setVisible(false);
    }
  };

  const handleHomeAddress = (address, key) => {
    localStorage.setItem("address", JSON.stringify({...address, id:key}));
    console.log(key);
    setHomeTag(key);
  };

  return (
    <>
      {loading ? (
        <Spinner title={"Setting up Your Profile :)"} />
      ) : (
        <>
          <div
            className="userpage-body"
            style={{ backgroundColor: "rgb(246,245,249)" }}
          >
            <Navbar inside={true} />
            <Card className="avatar-outerbox">
              <div className="avatar-flexbox">
                <label htmlFor="image-upload" className="image-upload-label">
                  <Avatar
                    image={
                      imagePreview || (!userData?.url ? user : userData.url)
                    }
                    size="xlarge"
                    className="profile-avatar"
                    shape="circle"
                    onMouseOver={showPopup}
                  />
                </label>
                <input
                  type="file"
                  id="image-upload"
                  hidden
                  accept="image/png, image/jpeg"
                  onChange={(e) => postDetails(e.target.files[0])}
                />

                <div>
                  <h1>{nameuser}</h1>
                  <p style={{ margin: "5px 0" }}>{emailuser}</p>
                </div>
              </div>
              {picture && (
                <Button
                  className="primereact-class"
                  raised
                  onClick={(e) => handleImage(e)}
                >
                  Save
                </Button>
              )}
              <h3 className="avatar-phone">{!phone ? "NA" : `+91-${phone}`}</h3>
              <Button
                icon="pi pi-sign-out"
                label="LogOut"
                severity="primary"
                className="primereact-class avatar-logout"
                onClick={Logout}
                raised
              />
            </Card>
            {/*Lower section*/}
            <TabView
              className="user-tabview"
              style={{ backgroundColor: "rgb(246,245,249)" }}
            >
              <TabPanel
                header="Orders"
                style={{ backgroundColor: "rgb(246,245,249)" }}
              >
                <div className="user-content">
                  <div className="user-orders-flexbox">
                    {order?.length > 0 &&
                      order.toReversed().map((singleOrder) => (
                        <Card
                          title={`Order ${singleOrder?.orderid}`}
                          className="user-orders-card"
                        >
                          <div className="user-order-card-content">
                            {singleOrder?.foods?.map((singleFood) => {
                              return (
                                <p className="order-food-text">
                                  <i
                                    style={{ margin: "0 1vh" }}
                                    className="pi pi-shopping-bag"
                                  ></i>
                                  {singleFood?.quantity} x{" "}
                                  {db[singleFood?.foodId - 1]?.title}
                                </p>
                              );
                            })}
                            <p className="order-food-text order-card-address">
                              <b>
                                {singleOrder?.totalPrice.toLocaleString(
                                  "en-US",
                                  { style: "currency", currency: "INR" }
                                )}
                              </b>
                            </p>
                            <div className="order-card-address">
                              <p className="order-card-address-text">
                                {singleOrder?.address?.address}
                              </p>
                              <p className="order-card-address-text">
                                {singleOrder?.address?.landmark}
                              </p>
                              <p className="order-card-address-text">
                                {singleOrder?.address?.pin}
                              </p>
                              <p className="order-card-address-text">
                                {singleOrder?.address?.contact}
                              </p>
                              <p className="order-card-address-text">
                                Placed on {singleOrder?.date} at{" "}
                                {singleOrder?.time}
                              </p>
                              <p className="order-card-address-text"></p>
                            </div>
                            <Button
                              style={{ float: "left" }}
                              className="primereact-class"
                              label=" Get Invoice"
                              onClick={() => {
                               singleOrder?.invoiceurl && window.open(singleOrder?.invoiceurl,'_blank') }
                                }
                            />
                            <Tag
                              severity="success"
                              value="Paid"
                              className="primereact-class"
                              style={{ float: "right" }}
                            ></Tag>
                          </div>
                        </Card>
                      ))}
                  </div>
                </div>
              </TabPanel>
              <TabPanel
                header="Address"
                style={{ backgroundColor: "rgb(246,245,249)" }}
              >
                <div className="user-content">
                  <Button
                    className="primereact-class"
                    label="Add"
                    icon="pi pi-external-link"
                    onClick={() => setVisible(true)}
                  ></Button>
                  <div className="user-content">
                    <Dialog
                      header="Address"
                      visible={visible}
                      style={{ width: "40vw" }}
                      onHide={() => {
                        if (!visible) return;
                        setVisible(false);
                      }}
                    >
                      <div className="profile-input-flexbox">
                        <label htmlFor="address">Address</label>
                        <InputTextarea
                          id="address"
                          autoResize
                          rows={3}
                          cols={30}
                          style={{ border: "2px solid rgb(238,240,242)" }}
                          className="primereact-class"
                          name="address"
                          onChange={handleAddress}
                          value={address?.address}
                        />
                      </div>
                      <div className="profile-input-flexbox">
                        <label htmlFor="landmark">Landmark</label>
                        <InputText
                          id="landmark"
                          aria-describedby="username-help"
                          className="primereact-class"
                          name="landmark"
                          value={address?.landmark}
                          onChange={handleAddress}
                        />
                      </div>
                      <div className="profile-input-flexbox">
                        <label htmlFor="pin">PIN</label>
                        <InputNumber
                          id="pin"
                          aria-describedby="username-help"
                          className="primereact-class"
                          value={!address?.pin ? "" : address?.pin}
                          name="pin"
                          onChange={(e) =>
                            setAddress((prev) => ({ ...prev, pin: e.value }))
                          }
                          useGrouping={false}
                        />
                      </div>

                      <Button
                        label="ADD"
                        className="profile-input-flexbox primereact-class"
                        outlined
                        onClick={handleAddressSave}
                      />
                    </Dialog>
                  </div>
                  <div className="user-orders-flexbox">
                    {addressRes?.length > 0 &&
                      addressRes.map((singleAddress, key) => {
                        return (
                          <div className="user-address-tag-card">
                            {homeTag == key && (
                              <Tag
                                severity="success"
                                value="Home"
                                style={{ float: "right", padding: "0.5vh 1vh" }}
                                rounded
                              ></Tag>
                            )}
                            <div
                              className="user-address-card"
                              onClick={() =>
                                handleHomeAddress(singleAddress, key)
                              }
                            >
                              <div className="order-card-address">
                                <p className="address-card-address-text">
                                  {singleAddress?.address}
                                </p>
                                <p className="address-card-address-text">
                                  {singleAddress?.landmark}
                                </p>
                                <p className="address-card-address-text">
                                  {singleAddress?.pin}
                                </p>
                              </div>
                            </div>
                          </div>
                        );
                      })}
                  </div>
                </div>
              </TabPanel>
              <TabPanel
                header="Account"
                style={{ backgroundColor: "rgb(246,245,249)" }}
              >
                <div className="user-content">
                  <div className="profile-input-outerbox">
                    <Card
                      style={{
                        width: "60%",
                        backgroundColor: "rgb(255,255,255)",
                      }}
                    >
                      <div className="profile-input-flexbox">
                        <label htmlFor="username">Name</label>
                        <InputText
                          id="username"
                          aria-describedby="username-help"
                          className="primereact-class"
                          value={profile?.name ? profile?.name : ""}
                          disabled={disableInput}
                          onChange={(e) => {
                            e.preventDefault();
                            setProfile((profile) => ({
                              ...profile,
                              name: e.target.value,
                            }));
                          }}
                        />
                      </div>
                      <div className="profile-input-flexbox">
                        <label htmlFor="username">Email</label>
                        <InputText
                          id="username"
                          aria-describedby="username-help"
                          className="primereact-class"
                          value={userData?.email}
                          disabled
                        />
                      </div>
                      <div className="profile-input-flexbox">
                        <label htmlFor="username">Phone</label>
                        <InputNumber
                          id="username"
                          aria-describedby="username-help"
                          className="primereact-class"
                          value={profile?.phone ? profile?.phone : ""}
                          disabled={disableInput}
                          onChange={(e) => {
                            setProfile((profile) => ({
                              ...profile,
                              phone: e.value,
                            }));
                          }}
                          useGrouping={false}
                        />
                      </div>
                      {disableInput ? (
                        <Button
                          label="Edit Profile"
                          className="profile-input-flexbox primereact-class"
                          outlined
                          onClick={() => {
                            setDisableInput(false);
                          }}
                        />
                      ) : (
                        <Button
                          label="Save Profile"
                          className="profile-input-flexbox primereact-class"
                          outlined
                          severity="success"
                          onClick={handleProfileSave}
                        />
                      )}
                    </Card>
                  </div>
                </div>
              </TabPanel>
            </TabView>
          </div>
        </>
      )}
    </>
  );
};

export default User;
