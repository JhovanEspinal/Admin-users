import { Component, OnInit,} from '@angular/core';
import { ApiService } from '../../service/api.service'
import { userListI } from 'src/app/models/userList.interface';
import {ActivatedRoute,Router} from '@angular/router'
import { AlertsService } from '../../serviceAlerts/alerts.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})


export class DashboardComponent implements OnInit {
  
  localUser : userListI ;
  anAdmin: boolean = false;
  users: userListI[] = []
  errorResult: boolean = false;

  constructor(private service: ApiService, private activateRoute : ActivatedRoute, private router: Router, private alertService: AlertsService) {

    this.localUser = {
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

    this.service.retrieveUser().subscribe(response => {
    this.users = response;
    this.activateRoute.params.subscribe(params =>{
    this.assignUser(params['id'])
    })
    })  
  }

  assignUser(id: string){
    this.users.forEach(user => {
      if( user._id === id ){
        this.localUser = user
      }
    });

    if( this.localUser.job === "administrador"){
      this.anAdmin = true
    }else{
      this.anAdmin = false;
    }

  }


  buscarUsuarios(termino: string) {
    this.service.searchUsers(termino).subscribe(response => {
      this.users = response;

      if (response.length === 0) {
        this.errorResult = true;
      } else {
        this.errorResult = false;
      }
    })
  }

  return(){
    this.activateRoute.params.subscribe(params =>{
      this.router.navigate(['login'])
  })
}

deleteUser(id : string, userName : string){
  
  this.service.deleteUser(id).subscribe(response =>{
        console.log("respuesta", response)
    if(response == null){
      this.alertService.showSucces("Operación Exitosa",`Se elimino el usuario ${userName}`)
      setTimeout(() => {
        window.location.reload(); 
       }, 1700);
    }
  },
  error => {
  this.alertService.showError('Error',`El usuario ${userName} no pudo ser eliminado`)
  }) 
 
}

editUser(id : string){
  let idLocalUser = this.localUser._id
  this.router.navigate(['edit',idLocalUser,id])
}




}
