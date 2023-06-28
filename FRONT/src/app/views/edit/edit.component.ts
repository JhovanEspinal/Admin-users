import { Component, OnInit } from '@angular/core';
import { FormGroup, Validators, FormBuilder} from '@angular/forms';
import {Router,ActivatedRoute} from '@angular/router'
import { userListI } from '../../models/userList.interface';
import { ApiService } from '../../service/api.service';
import { AlertsService } from '../../serviceAlerts/alerts.service';

@Component({
  selector: 'app-edit',
  templateUrl: './edit.component.html',
  styleUrls: ['./edit.component.css']
})
export class EditComponent implements OnInit {

  localUser : userListI;
  editForm : FormGroup;
  emailPattern: any = /^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/;
  passwordPattern: any = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&#.$($)$-$_])[A-Za-z\d$@$!%*?&#.$($)$-$_]{8,15}$/;


  constructor(private activateRoute: ActivatedRoute,private router: Router, private service : ApiService,private formBuilder: FormBuilder,private alert:AlertsService) { }

  ngOnInit(): void {

    let userId = this.activateRoute.snapshot.paramMap.get('user');
    this.service.retrieveUser().subscribe(response =>{
      response.forEach(user => {
        if(user._id === userId){
          this.localUser = user;
        }
      },
      error => {
       let errorMessage = <any>error;
       console.log(errorMessage);
      {
         this.alert.showError('Error', 'Error inesperado, intenta mas tarde')
       }
     });
    this.initializeForm(this.localUser)
    })
          
  }


  initializeForm(user : userListI){
  
    this.editForm = this.formBuilder.group({
      _id :        [user._id],
       name:       [user.name,[Validators.required,Validators.minLength(5),Validators.maxLength(20)]],
       cc:         [user.cc,[Validators.required,Validators.minLength(5)]],
       age:        [user.age, [Validators.required,Validators.min(18)]],
      gender:      [user.gender,Validators.required],
      job:         [user.job,Validators.required],
      description: [user.description,[Validators.required,Validators.maxLength(50)]],
      email:       [user.email,[Validators.required,Validators.pattern(this.emailPattern)]],
      password:    [user.password,[Validators.required,Validators.pattern(this.passwordPattern)]],
       img:        [user.img]
    })

   
  }


  return(){
    this.activateRoute.params.subscribe(params =>{
      this.router.navigate(['dashboard',params['id']])
  })
}


postForm(form: userListI){

  if(this.editForm.valid){

    if(form.gender === "otro"){
      form.img = "otro"
    }else if(form.gender != this.localUser.gender){
        let number = Math.floor(Math.random() * (5 - 1 + 1) + 1);
        let img = `${form.gender}${number}`
        form.img = img;
    }
 
     this.service.updateUser(form).subscribe(response =>{

    if(response == null){
      this.alert.showSucces('Operación exitosa',`Se guardaron los cambios del usuario ${form.name}`)
    }
   },

   error => {
    let errorMessage = <any>error;
    console.log(errorMessage);
   {
      this.alert.showError('Ups!', 'Ocurrio un error inesperado, intenta mas tarde')
    }
  })
}
}

get name(){ return this.editForm.get('name');}
get cc(){ return this.editForm.get('cc');}
get age(){ return this.editForm.get('age');}
get gender(){ return this.editForm.get('gender');}
get job(){ return this.editForm.get('job');}
get description(){ return this.editForm.get('description');}
get email(){ return this.editForm.get('email');}
get password(){ return this.editForm.get('password');}


}
