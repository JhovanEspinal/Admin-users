import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {LoginComponent } from './views/login/login.component';
import {NewComponent} from './views/new/new.component';
import {EditComponent} from './views/edit/edit.component';
import {DashboardComponent} from './views/dashboard/dashboard.component';



const routes: Routes = [
  {path:'',redirectTo:'login', pathMatch :'full'},
  {path: 'login', component:LoginComponent},
  {path: 'new/:id', component:NewComponent},
  {path: 'edit/:id/:user', component:EditComponent},
  {path: 'dashboard/:id', component:DashboardComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
export const routingComponents = [LoginComponent,DashboardComponent,NewComponent,EditComponent]