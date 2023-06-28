import { Component,OnInit } from '@angular/core';
import { FormGroup, Validators, FormBuilder} from '@angular/forms';
import {ActivatedRoute, Router} from '@angular/router';
import { userListI } from '../../models/userList.interface';
import { ApiService } from '../../service/api.service';
import { AlertsService } from '../../serviceAlerts/alerts.service';


@Component({
  selector: 'app-new',
  templateUrl: './new.component.html',
  styleUrls: ['./new.component.css']
})
export class NewComponent implements OnInit{

  userForm : FormGroup;
  emailPattern: any = /^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/;
  passwordPattern: any = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&#.$($)$-$_])[A-Za-z\d$@$!%*?&#.$($)$-$_]{8,15}$/;
  

  constructor(
    private router: Router,
    private activateRoute : ActivatedRoute,
     private service:ApiService, private alert:AlertsService,
      private formBuilder: FormBuilder){}
   

  ngOnInit():void{

  this.initializeForm();

  }

 
 initializeForm(){
  
  this.userForm = this.formBuilder.group({

     name:       ['',[Validators.required,Validators.minLength(5),Validators.maxLength(20)]],
     cc:         ['',[Validators.required,Validators.minLength(5)]],
     age:        ['', [Validators.required,Validators.min(18)]],
    gender:      ['',Validators.required],
    job:         ['',Validators.required],
    description: ['',[Validators.required,Validators.maxLength(50)]],
    email:       ['',[Validators.required,Validators.pattern(this.emailPattern)]],
    password:    ['',[Validators.required,Validators.pattern(this.passwordPattern)]],
     img:        [''],
  })
}


  return(){
    this.activateRoute.params.subscribe(params =>{
      this.router.navigate(['dashboard',params['id']])
  })
}


postForm(form: userListI){

  if(this.userForm.valid){

    if(form.gender === "otro"){
      form.img = "otro"
    }else{

      let number = Math.floor(Math.random() * (5 - 1 + 1) + 1);
      let img = `${form.gender}${number}`
      form.img = img;
    }
 
     this.service.saveUser(form).subscribe(response =>{

    if(response.email === form.email && response.password === form.password){
      this.alert.showSucces('Operación exitosa',`Se creo el usuario ${form.name}`)

    this.userForm.reset()
    }
   },

   error => {
    let errorMessage = <any>error;
    console.log(errorMessage);

    if (errorMessage.error == "REGISTERED_EMAIL\n"){
      this.alert.showError('Error', 'El correo electrónico ya se encuentra registrado')
    }
  })
}
}

get name(){ return this.userForm.get('name');}
get cc(){ return this.userForm.get('cc');}
get age(){ return this.userForm.get('age');}
get gender(){ return this.userForm.get('gender');}
get job(){ return this.userForm.get('job');}
get description(){ return this.userForm.get('description');}
get email(){ return this.userForm.get('email');}
get password(){ return this.userForm.get('password');}


}