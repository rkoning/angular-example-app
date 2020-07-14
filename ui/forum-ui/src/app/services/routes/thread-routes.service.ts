import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseRoutesService } from './base-routes.service';
import { Thread } from '@models/thread.model';

@Injectable({
  providedIn: 'root'
})
export class ThreadRoutesService extends BaseRoutesService<Thread> {
  public baseRoute = 'threads';
  constructor(protected http: HttpClient) { super(http); }
}
