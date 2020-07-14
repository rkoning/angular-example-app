import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseRoutesService } from './base-routes.service';
import { User } from '@models/user.model';

@Injectable({
  providedIn: 'root'
})
export class UserRoutesService extends BaseRoutesService<User> {
  public baseRoute = 'users';
  constructor(protected http: HttpClient) { super(http); }
}
