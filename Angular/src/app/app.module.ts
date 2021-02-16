import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { FormsModule } from '@angular/forms';
import { NgZorroAntdModule } from 'ng-zorro-antd'
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomeComponent } from './components/home/home.component';
import { PersonalComponent } from './components/personal/personal.component';
import { PersonalTimetableComponent } from './components/personal-timetable/personal-timetable.component';
import { ClassTimetableComponent } from './components/class-timetable/class-timetable.component';
import { GradesRecordComponent } from './components/grades-record/grades-record.component';
import { CourseSelectionComponent } from './components/course-selection/course-selection.component';
import { CourseArrangingComponent } from './components/course-arranging/course-arranging.component';
import { ShowCourseComponent } from './components/show-course/show-course.component';
import { ShowTeacherComponent } from './components/show-teacher/show-teacher.component';
import { ShowStudentComponent } from './components/show-student/show-student.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent,
    PersonalComponent,
    PersonalTimetableComponent,
    ClassTimetableComponent,
    GradesRecordComponent,
    CourseSelectionComponent,
    CourseArrangingComponent,
    ShowCourseComponent,
    ShowTeacherComponent,
    ShowStudentComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    NgZorroAntdModule,
    HttpClientModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
