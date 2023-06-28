import { Component, OnInit } from '@angular/core';
import{FormGroup,FormControl,Validators} from '@angular/forms'
import {ApiService} from '../../service/api.service'
import {ResponseI} from '../../models/response.interface'
import { LoginI, } from 'src/app/models/login.interface';
import {Router} from '@angular/router';
import { AlertsService } from '../../serviceAlerts/alerts.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})


export class LoginComponent implements OnInit {

  loginForm = new FormGroup({
    email: new FormControl('',Validators.required),
    password: new FormControl('',Validators.required)
  })

  public errorStatus : boolean = false;
  public errorMessage : string = "";
  public user : ResponseI


  constructor(private service :ApiService, private router: Router,private alert:AlertsService) { 
    this.user = {
      _id         : "",
      name        : "",
      cc          : 0,
      age         : 0,
      gender      : "",
      job         : "",
      description :"",
      email       : "",
      password    :"",
      img         :"",
  }

  }

  ngOnInit(): void {
  }

  onLogin(form: LoginI){

   this.service.loginUser(form).subscribe(response =>{

    if(response != null){
      this.alert.showSucces('Cool! ', 'Validación exitosa')
      this.user = response;
      this.errorMessage = "";
      this.errorStatus  = false;
      this.router.navigate(['dashboard',response._id])
    }
    },
    error => {
      let errorMessage = <any>error;
      console.log(errorMessage);

      if (errorMessage.error == "INVALID_USER_OR_PASSWORD\n"){
        this.errorStatus = true; 
          this.errorMessage = "Usuario o contraseña invalida";  
        console.log(this.errorMessage)
        
      }
    }) 
  }

  
removeError(){
  this.errorStatus = false;
}
}
