import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-class-timetable',
  templateUrl: './class-timetable.component.html',
  styleUrls: ['./class-timetable.component.scss']
})
export class ClassTimetableComponent implements OnInit {
  constructor(public http:HttpClient, 
    private router:Router,
    private notification: NzNotificationService) { }

  ngOnInit() {
    this.role = window.localStorage.getItem('role')
    this.getTable("")
  }

  getRole(){
    return this.role != "0"
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

  class = []
  selectedClass:string;
  role:string

  classChange(value: string): void {
    this.firstfirsts = []
    this.firstseconds = []
    this.firstthirds = []
    this.firstforths = []
    this.firstfifths = []
    this.secondfirsts = []
    this.secondseconds = []
    this.secondthirds = []
    this.secondforths = []
    this.secondfifths = []
    this.thirdfirsts = []
    this.thirdseconds = []
    this.thirdthirds = []
    this.thirdforths = []
    this.thirdfifths = []
    this.forthfirsts = []
    this.forthseconds = []
    this.forththirds = []
    this.forthforths = []
    this.forthfifths = []
    this.fifthfirsts = []
    this.fifthseconds = []
    this.fifththirds = []
    this.fifthforths = []
    this.fifthfifths = []
    this.selectedClass = value;
    this.getTable(this.selectedClass)
  }

  getTable(value : string){
    var account = window.localStorage.getItem('account')
    
    const httpOptions = {headers: new HttpHeaders({'Content-Type':'application/json'})};
    var api='http://127.0.0.1:8080/getClassTable';
    this.http.post(api, {'account':account, 'role':this.role, 'class':value}, httpOptions).subscribe((response:any)=>{
      if(response.check){
        if(this.role == "0"){
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
          if(!value) {
            this.class = response.class
            this.selectedClass = this.class[0]
          }
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
