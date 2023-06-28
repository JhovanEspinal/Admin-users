import { Injectable } from '@angular/core';
import { LoginI } from '../models/login.interface';
import { ResponseI } from '../models/response.interface'
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { Observable } from 'rxjs';
import { userListI } from '../models/userList.interface';
import { map } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class ApiService {

  url: string = "http://localhost:8080"


  constructor(private http: HttpClient) { }

  loginUser(form: LoginI): Observable<ResponseI> {

    let headers = new HttpHeaders().set('Content-Type', 'application/json');
    return this.http.post<ResponseI>(`${this.url}/user/validate`, form, { headers })
  }

  retrieveUser(): Observable<userListI[]> {

    let headers = new HttpHeaders().set('Content-Type', 'application/json');
    return this.http.get<userListI[]>(`${this.url}/user/retrieveUsers`, { headers })
  }

  saveUser(form:userListI): Observable<userListI>{
    let headers = new HttpHeaders().set('Content-Type', 'application/json');
    return this.http.post<userListI>(`${this.url}/user/save`, form, { headers })
  }
  
  updateUser(form:userListI): Observable<null>{
    let headers = new HttpHeaders().set('Content-Type', 'application/json');
    return this.http.put<null>(`${this.url}/user/update`, form, { headers })
  }

  deleteUser(id: string) :Observable<any>{
    let headers = new HttpHeaders().set('Content-Type', 'application/json');
    return this.http.delete<any>(`${this.url}/user/delete/${id}`,{headers})

  }


  searchUsers(termino: string): Observable<userListI[]> {
    let headers = new HttpHeaders().set('Content-Type', 'application/json');
    return this.http.get<userListI[]>(`${this.url}/user/retrieveUsers`, { headers }).pipe(
      
      map(( response: userListI[]) =>{ 
       const  users : userListI[] =[];
        termino = termino.toLowerCase();
        
        for (let index = 0; index < response.length; index++) {
          const nombre = response[index].name.toLowerCase();
          if (nombre.indexOf(termino) >=0){
            users.push(response[index])
          }
        } 

        return users;
      }
      )
    )

  }

  

}
