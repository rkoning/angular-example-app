import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseRoutesService } from './base-routes.service';
import { Thread } from '@models/thread.model';
import { Observable } from 'rxjs';
import { take } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ThreadRoutesService extends BaseRoutesService<Thread> {
  public baseRoute = 'threads';
  constructor(protected http: HttpClient) { super(http); }

  public getComments(id: string): Observable<Comment[]> {
    return this.http.get<Comment[]>(`${this.apiRoot}/${this.baseRoute}/comments`).pipe(take(1));
  }
}
