import React from "react";
import axios from "axios";

class Register extends React.Component{
    constructor(props) {
        super(props);

        this.submitRegister=this.submitRegister.bind(this);
    }

    submitRegister(e){
        e.preventDefault();
        var regexPhoneNumber=new RegExp("/[0-9]+/");
        var phoneNumber=document.getElementById("phone").value;
/*        if(!regexPhoneNumber.test(phoneNumber)){
            console.log("Format phone number is incorrect");
            console.log(phoneNumber)
            document.getElementById("error-register").innerHTML="Format phone number is incorrect";
            return;
        }*/
        var password=document.getElementById("password").value;
        if(password.length<8 || password.length>50){
            document.getElementById("error-register").innerHTML="Password must be longer than 8 and shorter than 50";
            return;
        }
        if(password.includes(" ")){
            document.getElementById("error-register").innerHTML="Cannot have white space in password";
            return;
        }
        var fullname=document.getElementById("fullname").value;
        if(fullname.length<8 || fullname.length>50){
            document.getElementById("error-register").innerHTML="Fullname must be longer than 8 and shorter than 50";
            return;
        }
        if(fullname.charAt(0)==' ' || fullname.charAt(-1)==' '){
            document.getElementById("error-register").innerHTML="Fullname cannot have space at the beginning or the end";
            return;
        }
        var register=`PhoneNumber=${phoneNumber}&Password=${password}&FullName=${fullname}`;
        axios.post(
            "http://127.0.0.1:9000/api/RegisterUser",
            register
        ).then(response=>{
            console.log(response);
        }).catch(err=>{
            console.log(err);
        })
    }

    render(){
        return (
            <div>
                <div id="error-register"></div>
                <form onSubmit={this.submitRegister}>
                    <label>Phone:</label>
                    <input type="text" id="phone" name="phone"/>
                    <label>Password:</label>
                    <input type="password" id="password" name="password"/>
                    <label>Fullname:</label>
                    <input type="text" id="fullname" name="fullname"/>
                    <input type="submit" value="Register"/>
                </form>

            </div>
        )
    }
}

export default Register;