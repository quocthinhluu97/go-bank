import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { AuthGuardGuard} from './services/guards/auth-guard.guard';
import { DashboardComponent } from './dashboard/dashboard.component';

const routes: Routes = [
    {path: '', component: LoginComponent},
    {path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuardGuard]},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }
