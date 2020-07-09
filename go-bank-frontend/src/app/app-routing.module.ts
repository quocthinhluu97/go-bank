import { NgModule  } from '@angular/core';
import { Routes, RouterModule  } from '@angular/router';
import { AuthGuardGuard  } from './services/guards/auth-guard.guard';
import { UserResolverService  } from './services/user-resolver/user-resolver.service';

import { LoginComponent  } from './login/login.component';
import { DashboardComponent  } from './dashboard/dashboard.component';
import { RegisterComponent  } from './register/register.component';
import {TransactionComponent} from './transaction/transaction.component';


const routes: Routes = [
    { path: 'login', component: LoginComponent  },

    { path: 'register', component: RegisterComponent  },

    { path: 'transaction',
      component: TransactionComponent,
      canActivate: [AuthGuardGuard],
      resolve: { user: UserResolverService }
    },

    {
        path: '',
        component: DashboardComponent,
        canActivate: [AuthGuardGuard],
        resolve: { user: UserResolverService  }
    }

];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]

})
export class AppRoutingModule {  }
