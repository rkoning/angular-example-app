import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { take } from 'rxjs/operators';
import { JsonResponse } from '@models/json-response.model';

@Injectable({
  providedIn: 'root'
})
export class BaseRoutesService<T> {

  public apiRoot = 'http://localhost:10000';
  public baseRoute: string;

  constructor(protected http: HttpClient) { }

  public get(itemId: string): Observable<JsonResponse<T>> {
    return this.http.get<JsonResponse<T>>(`${this.apiRoot}/${this.baseRoute}/${itemId}`).pipe(take(1));
  }

  public list(): Observable<JsonResponse<T[]>> {
    return this.http.get<JsonResponse<T[]>>(`${this.apiRoot}/${this.baseRoute}`).pipe(take(1));
  }

  public add(jsonBody: Partial<T>): Observable<JsonResponse<{InsertedID: string}>> {
    return this.http.post<JsonResponse<{ InsertedID: string }>>(`${this.apiRoot}/${this.baseRoute}`, jsonBody).pipe(take(1));
  }

  public update(jsonBody: Partial<T>): Observable<JsonResponse<T>> {
    return this.http.put<JsonResponse<T>>(`${this.apiRoot}/${this.baseRoute}`, jsonBody).pipe(take(1));
  }

  public delete(itemId: string): Observable<JsonResponse<T>> {
    return this.http.delete<JsonResponse<T>>(`${this.apiRoot}/${this.baseRoute}/${itemId}`).pipe(take(1));
  }
}
