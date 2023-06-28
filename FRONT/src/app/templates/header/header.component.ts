import { Component, EventEmitter, Input, Output, } from '@angular/core';
import { userListI } from '../../models/userList.interface';
import {Router} from '@angular/router'


@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent  {
  
@Input() userF : userListI; 
amAdmin :boolean = false;

@Output() public changeTermino: EventEmitter<string> = new EventEmitter()


constructor(private router: Router){
  this.userF = {
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

buscarUsuario(termino: any){
  this.changeTermino.emit(termino)
  }

newUser(){
  this.router.navigate(['new',this.userF._id])
} 

 


}
