CREATE TABLE "public"."registration"(member_id        integer,
                                     mobile           character varying(256),
                                     mobile_prefix     character varying(256),
                                     email            character varying(256),
                                     status           character varying(256),
                                     register_ip       character varying(256),
                                     account_status    character varying(256),
                                     createdat        integer ,
                                     register_platform character varying(256) 
                                     );

CREATE TABLE "public"."deposit"(id              bigint,
                                coin            character varying(256),
                                chain_detail     character varying(256),
                                member_id        integer,
                                amount          real NOT NULL,
                                status          integer,
                                request_time     integer,
                                request_deadline integer,
                                deposit_label    integer,
                                updated_time     integer
                               );


CREATE TABLE "public"."login"(member_id  integer,
                              login_time integer,
                              loginip    character varying(256),
                              platform   character varying(256));

CREATE TABLE "public"."logout"(member_id   integer,
                               logout_time integer,
                               logoutip    character varying(256),
                               platform    character varying(256));


CREATE TABLE "public"."withdrawals"(id          bigint,
                                    member_id    integer,
                                    coin        character varying(256),
                                    status      character varying(256),
                                    amount      character varying(256),
                                    fee         character varying(256),
                                    address     character varying(256),
                                    txid        character varying(256),
                                    submitedat  integer,
                                    successedat integer,
                                    updatedat   integer,
                                    chain_type   character varying(256),
                                    hash_info    character varying(256)
                                    );