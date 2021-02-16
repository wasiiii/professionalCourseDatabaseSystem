import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-show-teacher',
  templateUrl: './show-teacher.component.html',
  styleUrls: ['./show-teacher.component.scss']
})
export class ShowTeacherComponent implements OnInit {

  constructor(public http:HttpClient, 
    private router:Router,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.getAllTea()
  }

  tea
  classMap
  getAllTea(){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/getAllTea';
    this.http.post(api, {'account':account, 'role':role}, httpOptions).subscribe((response:any)=>{
      //console.log(response)
      if(response.check){
        this.tea = response.allTea
        this.classMap = response.class
      }
    });
  }

}
