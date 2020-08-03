import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseRoutesService } from './base-routes.service';
import { Thread } from '@models/thread.model';
import { Observable } from 'rxjs';
import { take } from 'rxjs/operators';
import { JsonResponse } from '@models/json-response.model';
import { Comment } from '@models/comment.model';

@Injectable({
  providedIn: 'root'
})
export class ThreadRoutesService extends BaseRoutesService<Thread> {
  public baseRoute = 'threads';
  constructor(protected http: HttpClient) { super(http); }

  public getComments(id: string): Observable<JsonResponse<Comment[]>> {
    return this.http.get<JsonResponse<Comment[]>>(`${this.apiRoot}/${this.baseRoute}/${id}/comments`).pipe(take(1));
  }

  public addComment(id: string, data: Comment): Observable<JsonResponse<{ insertedID: string }>> {
    return this.http.post<JsonResponse<{ insertedID: string }>>(`${this.apiRoot}/${this.baseRoute}/${id}/comments`, data)
      .pipe(take(1));
  }

  public deleteComment(threadId: string, commentId: string): Observable<JsonResponse<Comment>> {
    return this.http.delete<JsonResponse<Comment>>(`${this.apiRoot}/${this.baseRoute}/${threadId}/comments/${commentId}`).pipe(take(1));
  }
}
