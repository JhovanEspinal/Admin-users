import { Injectable } from '@angular/core';
import {ToastrService} from 'ngx-toastr'

@Injectable({
  providedIn: 'root'
})
export class AlertsService {

  constructor(private toast: ToastrService) { }


  showSucces(title : string, text :string){

    this.toast.success(text,title)
  }

  showError(title : string, text :string){

    this.toast.error(text,title)
  }


}
