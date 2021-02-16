import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-personal-timetable',
  templateUrl: './personal-timetable.component.html',
  styleUrls: ['./personal-timetable.component.scss']
})
export class PersonalTimetableComponent implements OnInit {

  constructor(public http:HttpClient, 
    private router:Router,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.getTable()
  }

  firstfirsts = []
  firstseconds = []
  firstthirds = []
  firstforths = []
  firstfifths = []
  secondfirsts = []
  secondseconds = []
  secondthirds = []
  secondforths = []
  secondfifths = []
  thirdfirsts = []
  thirdseconds = []
  thirdthirds = []
  thirdforths = []
  thirdfifths = []
  forthfirsts = []
  forthseconds = []
  forththirds = []
  forthforths = []
  forthfifths = []
  fifthfirsts = []
  fifthseconds = []
  fifththirds = []
  fifthforths = []
  fifthfifths = []
  showTea:boolean = false
  showClass:boolean = false

  getTable(){
    var account = window.localStorage.getItem('account')
    var role = window.localStorage.getItem('role')
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/getPersonalTable';
    this.http.post(api, {'account':account, 'role':role}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        if(role == "0"){
          this.showTea = true
          var table = response.table
          for (var i=0;i<table.length;i++){
            if (table[i].Section == 1 && table[i].Day == 1){
              this.firstfirsts.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 2){
              this.firstseconds.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 3){
              this.firstthirds.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 4){
              this.firstforths.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 5){
              this.firstfifths.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 1){
              this.secondfirsts.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 2){
              this.secondseconds.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 3){
              this.secondthirds.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 4){
              this.secondforths.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 5){
              this.secondfifths.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 1){
              this.thirdfirsts.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 2){
              this.thirdseconds.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 3){
              this.thirdthirds.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 4){
              this.thirdforths.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 5){
              this.thirdfifths.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 1){
              this.forthfirsts.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 2){
              this.forthseconds.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 3){
              this.forththirds.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 4){
              this.forthforths.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 5){
              this.forthfifths.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 1){
              this.fifthfirsts.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 2){
              this.fifthseconds.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 3){
              this.fifththirds.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 4){
              this.fifthforths.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 5){
              this.fifthfifths.push(table[i])
            }
          }
        }
        else{
          this.showClass = true
          var table = response.table
          for (var i=0;i<table.length;i++){
            if (table[i].Section == 1 && table[i].Day == 1){
              this.firstfirsts.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 2){
              this.firstseconds.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 3){
              this.firstthirds.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 4){
              this.firstforths.push(table[i])
            }else if (table[i].Section == 1 && table[i].Day == 5){
              this.firstfifths.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 1){
              this.secondfirsts.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 2){
              this.secondseconds.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 3){
              this.secondthirds.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 4){
              this.secondforths.push(table[i])
            }else if (table[i].Section == 2 && table[i].Day == 5){
              this.secondfifths.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 1){
              this.thirdfirsts.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 2){
              this.thirdseconds.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 3){
              this.thirdthirds.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 4){
              this.thirdforths.push(table[i])
            }else if (table[i].Section == 3 && table[i].Day == 5){
              this.thirdfifths.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 1){
              this.forthfirsts.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 2){
              this.forthseconds.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 3){
              this.forththirds.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 4){
              this.forthforths.push(table[i])
            }else if (table[i].Section == 4 && table[i].Day == 5){
              this.forthfifths.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 1){
              this.fifthfirsts.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 2){
              this.fifthseconds.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 3){
              this.fifththirds.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 4){
              this.fifthforths.push(table[i])
            }else if (table[i].Section == 5 && table[i].Day == 5){
              this.fifthfifths.push(table[i])
            }
          }
        }
      }
    });
  }

}
