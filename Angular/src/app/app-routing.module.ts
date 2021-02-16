import { ShowTeacherComponent } from './components/show-teacher/show-teacher.component';
import { ShowStudentComponent } from './components/show-student/show-student.component';
import { ShowCourseComponent } from './components/show-course/show-course.component';
import { PersonalTimetableComponent } from './components/personal-timetable/personal-timetable.component';
import { GradesRecordComponent } from './components/grades-record/grades-record.component';
import { CourseSelectionComponent } from './components/course-selection/course-selection.component';
import { CourseArrangingComponent } from './components/course-arranging/course-arranging.component';
import { ClassTimetableComponent } from './components/class-timetable/class-timetable.component';
import { LoginComponent } from './components/login/login.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard} from './auth-guard.service.service';
import { PersonalComponent } from './components/personal/personal.component';


const routes: Routes = [
  {
    path:'login',component:LoginComponent
  },
  {
    path:'personal',component:PersonalComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'class-timetable',component:ClassTimetableComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'course-arranging',component:CourseArrangingComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'course-selection',component:CourseSelectionComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'grades-record',component:GradesRecordComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'personal-timetable',component:PersonalTimetableComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'show-course',component:ShowCourseComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'show-student',component:ShowStudentComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'show-teacher',component:ShowTeacherComponent,
    canActivate:[AuthGuard],
  },
  {
    path:'**',redirectTo:'/personal',pathMatch:'full'
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers: [AuthGuard]
})
export class AppRoutingModule { }
