PGDMP  #    5                {        	   schoolApp    16.0    16.0     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16396 	   schoolApp    DATABASE     �   CREATE DATABASE "schoolApp" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE "schoolApp";
                postgres    false            �            1259    16410    class    TABLE     �   CREATE TABLE public.class (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    date_of_creation date NOT NULL
);
    DROP TABLE public.class;
       public         heap    postgres    false            �            1259    16409    class_id_seq    SEQUENCE     �   CREATE SEQUENCE public.class_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.class_id_seq;
       public          postgres    false    218            �           0    0    class_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.class_id_seq OWNED BY public.class.id;
          public          postgres    false    217            �            1259    16426 
   enrollment    TABLE     �   CREATE TABLE public.enrollment (
    student_id integer NOT NULL,
    class_id integer NOT NULL,
    date_of_enrollment date NOT NULL
);
    DROP TABLE public.enrollment;
       public         heap    postgres    false            �            1259    16401    students    TABLE     �   CREATE TABLE public.students (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    surname character varying(50) NOT NULL,
    email character varying(100)
);
    DROP TABLE public.students;
       public         heap    postgres    false            �            1259    16400    students_id_seq    SEQUENCE     �   CREATE SEQUENCE public.students_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.students_id_seq;
       public          postgres    false    216            �           0    0    students_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.students_id_seq OWNED BY public.students.id;
          public          postgres    false    215            $           2604    16413    class id    DEFAULT     d   ALTER TABLE ONLY public.class ALTER COLUMN id SET DEFAULT nextval('public.class_id_seq'::regclass);
 7   ALTER TABLE public.class ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            #           2604    16404    students id    DEFAULT     j   ALTER TABLE ONLY public.students ALTER COLUMN id SET DEFAULT nextval('public.students_id_seq'::regclass);
 :   ALTER TABLE public.students ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            �          0    16410    class 
   TABLE DATA           ;   COPY public.class (id, name, date_of_creation) FROM stdin;
    public          postgres    false    218   �       �          0    16426 
   enrollment 
   TABLE DATA           N   COPY public.enrollment (student_id, class_id, date_of_enrollment) FROM stdin;
    public          postgres    false    219   �       �          0    16401    students 
   TABLE DATA           <   COPY public.students (id, name, surname, email) FROM stdin;
    public          postgres    false    216   �       �           0    0    class_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.class_id_seq', 1, false);
          public          postgres    false    217            �           0    0    students_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.students_id_seq', 1, true);
          public          postgres    false    215            *           2606    16415    class class_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.class
    ADD CONSTRAINT class_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.class DROP CONSTRAINT class_pkey;
       public            postgres    false    218            ,           2606    16430    enrollment enrollment_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.enrollment
    ADD CONSTRAINT enrollment_pkey PRIMARY KEY (student_id, class_id);
 D   ALTER TABLE ONLY public.enrollment DROP CONSTRAINT enrollment_pkey;
       public            postgres    false    219    219            &           2606    16408    students students_email_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_email_key UNIQUE (email);
 E   ALTER TABLE ONLY public.students DROP CONSTRAINT students_email_key;
       public            postgres    false    216            (           2606    16406    students students_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.students DROP CONSTRAINT students_pkey;
       public            postgres    false    216            -           2606    16436 #   enrollment enrollment_class_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.enrollment
    ADD CONSTRAINT enrollment_class_id_fkey FOREIGN KEY (class_id) REFERENCES public.class(id);
 M   ALTER TABLE ONLY public.enrollment DROP CONSTRAINT enrollment_class_id_fkey;
       public          postgres    false    4650    219    218            .           2606    16431 %   enrollment enrollment_student_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.enrollment
    ADD CONSTRAINT enrollment_student_id_fkey FOREIGN KEY (student_id) REFERENCES public.students(id);
 O   ALTER TABLE ONLY public.enrollment DROP CONSTRAINT enrollment_student_id_fkey;
       public          postgres    false    219    216    4648            �      x������ � �      �      x������ � �      �   +   x�3��J�K�t�O��2�R�SR+srR���s�b���� �2     