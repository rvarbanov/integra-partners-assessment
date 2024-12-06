import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

export interface User {
  user_id: number;
  user_name: string;
  first_name: string;
  last_name: string;
  email: string;
  user_status: string;
  department: string;
  created_at?: string;
  updated_at?: string;
}

interface ApiResponse<T> {
  data: T;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  getUsers(): Observable<User[]> {
    return this.http.get<ApiResponse<User[]>>(`${this.apiUrl}/users`)
      .pipe(
        map(response => response.data),
        catchError(this.handleError)
      );
  }

  getUser(id: number): Observable<User> {
    return this.http.get<User>(`${this.apiUrl}/user/${id}`)
      .pipe(catchError(this.handleError));
  }

  addUser(user: User): Observable<User> {
    return this.http.post<User>(`${this.apiUrl}/user`, user)
      .pipe(catchError(this.handleError));
  }

  updateUser(user: User): Observable<User> {
    return this.http.put<User>(`${this.apiUrl}/user/${user.user_id}`, user)
      .pipe(catchError(this.handleError));
  }

  deleteUser(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/user/${id}`)
      .pipe(catchError(this.handleError));
  }

  private handleError(error: HttpErrorResponse) {
    let errorMessage = 'An error occurred';
    
    if (error.error instanceof ErrorEvent) {
      // Client-side error
      errorMessage = error.error.message;
    } else {
      // Server-side error
      if (error.status === 409) {
        errorMessage = 'Username already exists';
      } else if (error.error?.message) {
        errorMessage = error.error.message;
      } else {
        errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
      }
    }
    
    console.error('Error:', errorMessage);
    return throwError(() => new Error(errorMessage));
  }
} 