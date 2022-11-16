package main

// #include <stdio.h>
// #include <stdlib.h>
// #include <math.h>
/*
typedef unsigned char	BitSequence;

typedef struct _testParameters {
	int		n;
	int		blockFrequencyBlockLength;
	int		nonOverlappingTemplateBlockLength;
	int		overlappingTemplateBlockLength;
	int		serialBlockLength;
	int		linearComplexitySequenceLength;
	int		approximateEntropyBlockLength;
	int		numOfBitStreams;
} TP;



static void drfti1(int n, double* wa, int* ifac)
{
	static int		ntryh[4] = { 4,2,3,5 };
	static double	tpi = 6.28318530717958647692528676655900577;
	double			arg, argh, argld, fi;
	int				ntry = 0, i, j = -1;
	int				k1, l1, l2, ib;
	int				ld, ii, ip, is, nq, nr;
	int				ido, ipm, nfm1;
	int				nl = n;
	int				nf = 0;

L101:
	j++;
	if (j < 4)
		ntry = ntryh[j];
	else
		ntry += 2;

L104:
	nq = nl / ntry;
	nr = nl - ntry * nq;
	if (nr != 0)
		goto L101;

	nf++;
	ifac[nf + 1] = ntry;
	nl = nq;
	if (ntry != 2)
		goto L107;
	if (nf == 1)
		goto L107;

	for (i = 1; i < nf; i++) {
		ib = nf - i + 1;
		ifac[ib + 1] = ifac[ib];
	}
	ifac[2] = 2;

L107:
	if (nl != 1)
		goto L104;
	ifac[0] = n;
	ifac[1] = nf;
	argh = tpi / n;
	is = 0;
	nfm1 = nf - 1;
	l1 = 1;

	if (nfm1 == 0)
		return;

	for (k1 = 0; k1 < nfm1; k1++) {
		ip = ifac[k1 + 2];
		ld = 0;
		l2 = l1 * ip;
		ido = n / l2;
		ipm = ip - 1;

		for (j = 0; j < ipm; j++) {
			ld += l1;
			i = is;
			argld = (double)ld * argh;
			fi = 0.0;
			for (ii = 2; ii < ido; ii += 2) {
				fi += 1.0;
				arg = fi * argld;
				wa[i++] = cos(arg);
				wa[i++] = sin(arg);
			}
			is += ido;
		}
		l1 = l2;
	}
}


void __ogg_fdrffti(int n, double* wsave, int* ifac)
{
	if (n == 1)
		return;
	drfti1(n, wsave + n, ifac);
}



static void dradf2(int ido, int l1, double* cc, double* ch, double* wa1)
{
	int		i, k;
	double	ti2, tr2;
	int		t0, t1, t2, t3, t4, t5, t6;

	t1 = 0;
	t0 = (t2 = l1 * ido);
	t3 = ido << 1;
	for (k = 0; k < l1; k++) {
		ch[t1 << 1] = cc[t1] + cc[t2];
		ch[(t1 << 1) + t3 - 1] = cc[t1] - cc[t2];
		t1 += ido;
		t2 += ido;
	}

	if (ido < 2)
		return;
	if (ido == 2)
		goto L105;

	t1 = 0;
	t2 = t0;
	for (k = 0; k < l1; k++) {
		t3 = t2;
		t4 = (t1 << 1) + (ido << 1);
		t5 = t1;
		t6 = t1 + t1;
		for (i = 2; i < ido; i += 2) {
			t3 += 2;
			t4 -= 2;
			t5 += 2;
			t6 += 2;
			tr2 = wa1[i - 2] * cc[t3 - 1] + wa1[i - 1] * cc[t3];
			ti2 = wa1[i - 2] * cc[t3] - wa1[i - 1] * cc[t3 - 1];
			ch[t6] = cc[t5] + ti2;
			ch[t4] = ti2 - cc[t5];
			ch[t6 - 1] = cc[t5 - 1] + tr2;
			ch[t4 - 1] = cc[t5 - 1] - tr2;
		}
		t1 += ido;
		t2 += ido;
	}

	if (ido % 2 == 1)
		return;

L105:
	t3 = (t2 = (t1 = ido) - 1);
	t2 += t0;
	for (k = 0; k < l1; k++) {
		ch[t1] = -cc[t2];
		ch[t1 - 1] = cc[t3];
		t1 += ido << 1;
		t2 += ido;
		t3 += ido;
	}
}

static void dradf4(int ido, int l1, double* cc, double* ch, double* wa1,
	double* wa2, double* wa3)
{
	static double	hsqt2 = .70710678118654752440084436210485;
	int				i, k, t0, t1, t2, t3, t4, t5, t6;
	double			ci2, ci3, ci4, cr2, cr3, cr4;
	double			ti1, ti2, ti3, ti4, tr1, tr2, tr3, tr4;

	t0 = l1 * ido;
	t1 = t0;
	t4 = t1 << 1;
	t2 = t1 + (t1 << 1);
	t3 = 0;

	for (k = 0; k < l1; k++) {
		tr1 = cc[t1] + cc[t2];
		tr2 = cc[t3] + cc[t4];
		ch[t5 = t3 << 2] = tr1 + tr2;
		ch[(ido << 2) + t5 - 1] = tr2 - tr1;
		ch[(t5 += (ido << 1)) - 1] = cc[t3] - cc[t4];
		ch[t5] = cc[t2] - cc[t1];

		t1 += ido;
		t2 += ido;
		t3 += ido;
		t4 += ido;
	}

	if (ido < 2)
		return;
	if (ido == 2)
		goto L105;

	t1 = 0;
	for (k = 0; k < l1; k++) {
		t2 = t1;
		t4 = t1 << 2;
		t5 = (t6 = ido << 1) + t4;
		for (i = 2; i < ido; i += 2) {
			t3 = (t2 += 2);
			t4 += 2;
			t5 -= 2;

			t3 += t0;
			cr2 = wa1[i - 2] * cc[t3 - 1] + wa1[i - 1] * cc[t3];
			ci2 = wa1[i - 2] * cc[t3] - wa1[i - 1] * cc[t3 - 1];
			t3 += t0;
			cr3 = wa2[i - 2] * cc[t3 - 1] + wa2[i - 1] * cc[t3];
			ci3 = wa2[i - 2] * cc[t3] - wa2[i - 1] * cc[t3 - 1];
			t3 += t0;
			cr4 = wa3[i - 2] * cc[t3 - 1] + wa3[i - 1] * cc[t3];
			ci4 = wa3[i - 2] * cc[t3] - wa3[i - 1] * cc[t3 - 1];

			tr1 = cr2 + cr4;
			tr4 = cr4 - cr2;
			ti1 = ci2 + ci4;
			ti4 = ci2 - ci4;
			ti2 = cc[t2] + ci3;
			ti3 = cc[t2] - ci3;
			tr2 = cc[t2 - 1] + cr3;
			tr3 = cc[t2 - 1] - cr3;


			ch[t4 - 1] = tr1 + tr2;
			ch[t4] = ti1 + ti2;

			ch[t5 - 1] = tr3 - ti4;
			ch[t5] = tr4 - ti3;

			ch[t4 + t6 - 1] = ti4 + tr3;
			ch[t4 + t6] = tr4 + ti3;

			ch[t5 + t6 - 1] = tr2 - tr1;
			ch[t5 + t6] = ti1 - ti2;
		}
		t1 += ido;
	}
	if (ido % 2 == 1)
		return;

L105:

	t2 = (t1 = t0 + ido - 1) + (t0 << 1);
	t3 = ido << 2;
	t4 = ido;
	t5 = ido << 1;
	t6 = ido;

	for (k = 0; k < l1; k++) {
		ti1 = -hsqt2 * (cc[t1] + cc[t2]);
		tr1 = hsqt2 * (cc[t1] - cc[t2]);
		ch[t4 - 1] = tr1 + cc[t6 - 1];
		ch[t4 + t5 - 1] = cc[t6 - 1] - tr1;
		ch[t4] = ti1 - cc[t1 + t0];
		ch[t4 + t5] = ti1 + cc[t1 + t0];
		t1 += ido;
		t2 += ido;
		t4 += t3;
		t6 += ido;
	}
}

static void dradfg(int ido, int ip, int l1, int idl1, double* cc, double* c1,
	double* c2, double* ch, double* ch2, double* wa)
{
	static double	tpi = 6.28318530717958647692528676655900577;
	int				idij, ipph, i, j, k, l, ic, ik, is;
	int				t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10;
	double			dc2, ai1, ai2, ar1, ar2, ds2;
	int				nbd;
	double			dcp, arg, dsp, ar1h, ar2h;
	int				idp2, ipp2;

	arg = tpi / (double)ip;
	dcp = cos(arg);
	dsp = sin(arg);
	ipph = (ip + 1) >> 1;
	ipp2 = ip;
	idp2 = ido;
	nbd = (ido - 1) >> 1;
	t0 = l1 * ido;
	t10 = ip * ido;

	if (ido == 1)
		goto L119;
	for (ik = 0; ik < idl1; ik++)
		ch2[ik] = c2[ik];

	t1 = 0;
	for (j = 1; j < ip; j++) {
		t1 += t0;
		t2 = t1;
		for (k = 0; k < l1; k++) {
			ch[t2] = c1[t2];
			t2 += ido;
		}
	}

	is = -ido;
	t1 = 0;
	if (nbd > l1) {
		for (j = 1; j < ip; j++) {
			t1 += t0;
			is += ido;
			t2 = -ido + t1;
			for (k = 0; k < l1; k++) {
				idij = is - 1;
				t2 += ido;
				t3 = t2;
				for (i = 2; i < ido; i += 2) {
					idij += 2;
					t3 += 2;
					ch[t3 - 1] = wa[idij - 1] * c1[t3 - 1] + wa[idij] * c1[t3];
					ch[t3] = wa[idij - 1] * c1[t3] - wa[idij] * c1[t3 - 1];
				}
			}
		}
	}
	else {
		for (j = 1; j < ip; j++) {
			is += ido;
			idij = is - 1;
			t1 += t0;
			t2 = t1;
			for (i = 2; i < ido; i += 2) {
				idij += 2;
				t2 += 2;
				t3 = t2;
				for (k = 0; k < l1; k++) {
					ch[t3 - 1] = wa[idij - 1] * c1[t3 - 1] + wa[idij] * c1[t3];
					ch[t3] = wa[idij - 1] * c1[t3] - wa[idij] * c1[t3 - 1];
					t3 += ido;
				}
			}
		}
	}

	t1 = 0;
	t2 = ipp2 * t0;
	if (nbd < l1) {
		for (j = 1; j < ipph; j++) {
			t1 += t0;
			t2 -= t0;
			t3 = t1;
			t4 = t2;
			for (i = 2; i < ido; i += 2) {
				t3 += 2;
				t4 += 2;
				t5 = t3 - ido;
				t6 = t4 - ido;
				for (k = 0; k < l1; k++) {
					t5 += ido;
					t6 += ido;
					c1[t5 - 1] = ch[t5 - 1] + ch[t6 - 1];
					c1[t6 - 1] = ch[t5] - ch[t6];
					c1[t5] = ch[t5] + ch[t6];
					c1[t6] = ch[t6 - 1] - ch[t5 - 1];
				}
			}
		}
	}
	else {
		for (j = 1; j < ipph; j++) {
			t1 += t0;
			t2 -= t0;
			t3 = t1;
			t4 = t2;
			for (k = 0; k < l1; k++) {
				t5 = t3;
				t6 = t4;
				for (i = 2; i < ido; i += 2) {
					t5 += 2;
					t6 += 2;
					c1[t5 - 1] = ch[t5 - 1] + ch[t6 - 1];
					c1[t6 - 1] = ch[t5] - ch[t6];
					c1[t5] = ch[t5] + ch[t6];
					c1[t6] = ch[t6 - 1] - ch[t5 - 1];
				}
				t3 += ido;
				t4 += ido;
			}
		}
	}

L119:
	for (ik = 0; ik < idl1; ik++)
		c2[ik] = ch2[ik];

	t1 = 0;
	t2 = ipp2 * idl1;
	for (j = 1; j < ipph; j++) {
		t1 += t0;
		t2 -= t0;
		t3 = t1 - ido;
		t4 = t2 - ido;
		for (k = 0; k < l1; k++) {
			t3 += ido;
			t4 += ido;
			c1[t3] = ch[t3] + ch[t4];
			c1[t4] = ch[t4] - ch[t3];
		}
	}

	ar1 = 1.0;
	ai1 = 0.0;
	t1 = 0;
	t2 = ipp2 * idl1;
	t3 = (ip - 1) * idl1;
	for (l = 1; l < ipph; l++) {
		t1 += idl1;
		t2 -= idl1;
		ar1h = dcp * ar1 - dsp * ai1;
		ai1 = dcp * ai1 + dsp * ar1;
		ar1 = ar1h;
		t4 = t1;
		t5 = t2;
		t6 = t3;
		t7 = idl1;

		for (ik = 0; ik < idl1; ik++) {
			ch2[t4++] = c2[ik] + ar1 * c2[t7++];
			ch2[t5++] = ai1 * c2[t6++];
		}

		dc2 = ar1;
		ds2 = ai1;
		ar2 = ar1;
		ai2 = ai1;

		t4 = idl1;
		t5 = (ipp2 - 1) * idl1;
		for (j = 2; j < ipph; j++) {
			t4 += idl1;
			t5 -= idl1;

			ar2h = dc2 * ar2 - ds2 * ai2;
			ai2 = dc2 * ai2 + ds2 * ar2;
			ar2 = ar2h;

			t6 = t1;
			t7 = t2;
			t8 = t4;
			t9 = t5;
			for (ik = 0; ik < idl1; ik++) {
				ch2[t6++] += ar2 * c2[t8++];
				ch2[t7++] += ai2 * c2[t9++];
			}
		}
	}

	t1 = 0;
	for (j = 1; j < ipph; j++) {
		t1 += idl1;
		t2 = t1;
		for (ik = 0; ik < idl1; ik++)
			ch2[ik] += c2[t2++];
	}

	if (ido < l1)
		goto L132;

	t1 = 0;
	t2 = 0;
	for (k = 0; k < l1; k++) {
		t3 = t1;
		t4 = t2;
		for (i = 0; i < ido; i++)
			cc[t4++] = ch[t3++];
		t1 += ido;
		t2 += t10;
	}

	goto L135;

L132:
	for (i = 0; i < ido; i++) {
		t1 = i;
		t2 = i;
		for (k = 0; k < l1; k++) {
			cc[t2] = ch[t1];
			t1 += ido;
			t2 += t10;
		}
	}

L135:
	t1 = 0;
	t2 = ido << 1;
	t3 = 0;
	t4 = ipp2 * t0;
	for (j = 1; j < ipph; j++) {

		t1 += t2;
		t3 += t0;
		t4 -= t0;

		t5 = t1;
		t6 = t3;
		t7 = t4;

		for (k = 0; k < l1; k++) {
			cc[t5 - 1] = ch[t6];
			cc[t5] = ch[t7];
			t5 += t10;
			t6 += ido;
			t7 += ido;
		}
	}

	if (ido == 1)
		return;
	if (nbd < l1)
		goto L141;

	t1 = -ido;
	t3 = 0;
	t4 = 0;
	t5 = ipp2 * t0;
	for (j = 1; j < ipph; j++) {
		t1 += t2;
		t3 += t2;
		t4 += t0;
		t5 -= t0;
		t6 = t1;
		t7 = t3;
		t8 = t4;
		t9 = t5;
		for (k = 0; k < l1; k++) {
			for (i = 2; i < ido; i += 2) {
				ic = idp2 - i;
				cc[i + t7 - 1] = ch[i + t8 - 1] + ch[i + t9 - 1];
				cc[ic + t6 - 1] = ch[i + t8 - 1] - ch[i + t9 - 1];
				cc[i + t7] = ch[i + t8] + ch[i + t9];
				cc[ic + t6] = ch[i + t9] - ch[i + t8];
			}
			t6 += t10;
			t7 += t10;
			t8 += ido;
			t9 += ido;
		}
	}
	return;

L141:

	t1 = -ido;
	t3 = 0;
	t4 = 0;
	t5 = ipp2 * t0;
	for (j = 1; j < ipph; j++) {
		t1 += t2;
		t3 += t2;
		t4 += t0;
		t5 -= t0;
		for (i = 2; i < ido; i += 2) {
			t6 = idp2 + t1 - i;
			t7 = i + t3;
			t8 = i + t4;
			t9 = i + t5;
			for (k = 0; k < l1; k++) {
				cc[t7 - 1] = ch[t8 - 1] + ch[t9 - 1];
				cc[t6 - 1] = ch[t8 - 1] - ch[t9 - 1];
				cc[t7] = ch[t8] + ch[t9];
				cc[t6] = ch[t9] - ch[t8];
				t6 += t10;
				t7 += t10;
				t8 += ido;
				t9 += ido;
			}
		}
	}
}


static void drftf1(int n, double* c, double* ch, double* wa, int* ifac)
{
	int		i, k1, l1, l2;
	int		na, kh, nf;
	int		ip, iw, ido, idl1, ix2, ix3;

	nf = ifac[1];
	na = 1;
	l2 = n;
	iw = n;

	for (k1 = 0; k1 < nf; k1++) {
		kh = nf - k1;
		ip = ifac[kh + 1];
		l1 = l2 / ip;
		ido = n / l2;
		idl1 = ido * l1;
		iw -= (ip - 1) * ido;
		na = 1 - na;

		if (ip != 4)
			goto L102;

		ix2 = iw + ido;
		ix3 = ix2 + ido;
		if (na != 0)
			dradf4(ido, l1, ch, c, wa + iw - 1, wa + ix2 - 1, wa + ix3 - 1);
		else
			dradf4(ido, l1, c, ch, wa + iw - 1, wa + ix2 - 1, wa + ix3 - 1);
		goto L110;

	L102:
		if (ip != 2)
			goto L104;
		if (na != 0)
			goto L103;

		dradf2(ido, l1, c, ch, wa + iw - 1);
		goto L110;

	L103:
		dradf2(ido, l1, ch, c, wa + iw - 1);
		goto L110;

	L104:
		if (ido == 1)
			na = 1 - na;
		if (na != 0)
			goto L109;

		dradfg(ido, ip, l1, idl1, c, c, c, ch, ch, wa + iw - 1);
		na = 1;
		goto L110;

	L109:
		dradfg(ido, ip, l1, idl1, ch, ch, ch, c, c, wa + iw - 1);
		na = 0;

	L110:
		l2 = l1;
	}

	if (na == 1)
		return;

	for (i = 0; i < n; i++)
		c[i] = ch[i];
}


void __ogg_fdrfftf_ex(int n, double* r, double* wsave, int* ifac)
{
	if (n == 1)
		return;
	drftf1(n, r, wsave, wsave + n, ifac);
}


double nist_erfc(double x)
{
	double t, z, ans, result;
	z = fabs(x);
	t = 1.0 / (1.0 + 0.5 * z);

	ans = t * exp(-z * z - 1.26551223 + t * (1.00002368 + t * (0.37409196 + t * (0.09678418 +
		t * (-0.18628806 + t * (0.27886807 + t * (-1.13520398 + t * (1.48851587 +
			t * (-0.82215223 + t * 0.17087277)))))))));

	if (x >= 0.0) {
		result = ans;
	}
	else {
		result = 2.0 - ans;
	}
	return result;
}

static unsigned char _compute(unsigned char b, unsigned char factor)
{
	if ((factor & b) == factor) {
		return 0x01;
	}
	else {
		return 0x00;
	}
}

int BytesToBitSequence(unsigned char* in, int inl, BitSequence* outbuf, int bufsize)
{
	int j = 0, i = 0;
	if (bufsize < inl * 8) {
		return 0;
	}

	for (i = 0; i < inl; ++i) {
		j = i * 8;
		outbuf[j] = (BitSequence)(_compute(in[i], 0x80));
		outbuf[j + 1] = (BitSequence)(_compute(in[i], 0x40));
		outbuf[j + 2] = (BitSequence)(_compute(in[i], 0x20));
		outbuf[j + 3] = (BitSequence)(_compute(in[i], 0x10));
		outbuf[j + 4] = (BitSequence)(_compute(in[i], 0x08));
		outbuf[j + 5] = (BitSequence)(_compute(in[i], 0x04));
		outbuf[j + 6] = (BitSequence)(_compute(in[i], 0x02));
		outbuf[j + 7] = (BitSequence)(_compute(in[i], 0x01));
	}
	return 1;
}



int myDiscreteFourierTransform(unsigned char* in,int n,double *p_value,double *q_value)
{
	BitSequence* epsilon = (BitSequence*)calloc(n, sizeof(BitSequence));
	if(!epsilon)
		return 0.0;
	BytesToBitSequence(in,n/8,epsilon,n);

	double	upperBound, percentile, N_l, N_o, V, * m=NULL, * X, * wsave = NULL;

	int		i, count, ifac[15];

	if (((X = (double*)calloc(n, sizeof(double))) == ((void*)0)) ||
		((wsave = (double*)calloc(2 * n, sizeof(double))) == ((void*)0)) ||
		((m = (double*)calloc(n / 2 + 1, sizeof(double))) == ((void*)0))) {
		if (X != ((void*)0))
			free(X);
		if (wsave != ((void*)0))
			free(wsave);
		if (m != ((void*)0))
			free(m);

		return 0;
	}
	for (i = 0; i < n; i++)
		X[i] = 2 * (int)epsilon[i] - 1;

	__ogg_fdrffti(n, wsave, ifac);

	__ogg_fdrfftf_ex(n, X, wsave, ifac);

	m[0] = sqrt(X[0] * X[0]);

	for (i = 0; i < n / 2; i++)
		m[i + 1] = sqrt(pow(X[2 * i + 1], 2) + pow(X[2 * i + 2], 2));
	count = 0;
	upperBound = sqrt(2.995732274 * n);

	for (i = 0; i < n / 2; i++)
	{
		if (fabs(m[i]) < upperBound)
			count++;
	}

	N_l = (double)count;
	N_o = (double)0.95 * n / 2.0;

	V = (N_l - N_o) / sqrt(n / 3.8 * 0.95 * 0.05);
	*p_value = nist_erfc(fabs(V) / sqrt(2.0));

	*q_value = nist_erfc(V / sqrt(2.0))/2;

	free(X);
	free(wsave);
	free(m);
	free(epsilon);

	return 0;
}

*/
import "C"
import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/Trisia/randomness"
	"github.com/Trisia/randomness/detect"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func fun_DiscreteFourierTransformTest(buf []byte, bitLen int) (float64, float64) {

	var P C.double = 0.0

	var Q C.double = 0.0

	buffer := unsafe.Pointer(&buf[0])

	C.myDiscreteFourierTransform((*C.uchar)(buffer), C.int(bitLen), &P, &Q)

	return float64(P), float64(Q)
}

var mainwin *ui.Window

var itemStr = [27]string{"单比特频数检测", "块内频数检测 m=10000", "扑克检测 m=4", "扑克检测 m=8", "重叠子序列检测 m=3 P1", "重叠子序列检测 m=3 P2", "重叠子序列检测 m=5 P1", "重叠子序列检测 m=5 P2", "游程总数检测", "游程分布检测", "块内最大1游程检测 m=10000", "块内最大0游程检测 m=10000", "二元推导检测 k=3", "二元推导检测 k=7", "自相关检测 d=1", "自相关检测 d=2", "自相关检测 d=8", "自相关检测 d=16", "矩阵秩检测", "累加和前向检测", "累加和后向检测", "近似熵检测 m=2", "近似熵检测 m=5", "线性复杂度检测 m=500", "线性复杂度检测 m=1000", "Maurer通用统计检测 L=7 Q=1280", "离散傅里叶检测"}

var selectArry [len(itemStr)]int
var selectslice []int
var cnt []int32

var checkboxs [len(itemStr)]*ui.Checkbox
var labels_succ [len(itemStr)]*ui.Label
var labels_fail [len(itemStr)]*ui.Label

var failcnt = make([]int32, len(itemStr))

var labelsetbit *ui.Label
var labelsetNum *ui.Label
var labelsetAlphaNum *ui.Label
var labeltime *ui.Label
var labeltimebegin *ui.Label
var labeltimeend *ui.Label
var buttonOK *ui.Button
var buttonFile *ui.Button

var setcnt int

// Alpha 显著性水平α
// AlphaT 分布均匀性的显著性水平
// const AlphaT float64 = 0.0001
var Alpha float64
var entryAlpha *ui.Entry

var entryPath *ui.Entry
var entryOutPath *ui.Entry

var reportPath string
var inputPath string

var prosessbar *ui.ProgressBar

var seletctAll bool

type R struct {
	Name string
	P    []float64
	Q    []float64
}

var Perr [len(itemStr)]float64
var Qerr [len(itemStr)]float64

func ROUNDOverlapping(m int, bits []bool, wg *sync.WaitGroup) {
	defer wg.Done()

	p1, p2, q1, q2 := randomness.OverlappingTemplateMatchingProto(bits, m)
	if 3 == m {
		Perr[4] = p1
		Perr[5] = p2
		Qerr[4] = q1
		Qerr[5] = q2

		//fmt.Printf("index=4 %s\n", itemStr[4])
		//fmt.Printf("index=5 %s\n", itemStr[5])
		ui.QueueMain(func() {
			labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
			prosessbar.SetValue(int(math.Min(float64(prosessbar.Value()+4), 97)))
			if Perr[4] >= 0.01 {
				labels_succ[4].SetText("1")
			} else {
				labels_fail[4].SetText("1")
			}
			if Perr[5] >= 0.01 {
				labels_succ[5].SetText("1")
			} else {
				labels_fail[5].SetText("1")
			}
		})
	}

	if 5 == m {
		Perr[6] = p1
		Perr[7] = p2
		Qerr[6] = q1
		Qerr[7] = q2
		//fmt.Printf("index=6 %s\n", itemStr[6])
		//fmt.Printf("index=7 %s\n", itemStr[7])

		ui.QueueMain(func() {
			labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
			prosessbar.SetValue(int(math.Min(float64(prosessbar.Value()+4), 97)))
			if Perr[6] >= 0.01 {
				labels_succ[6].SetText("1")
			} else {
				labels_fail[6].SetText("1")
			}
			if Perr[7] >= 0.01 {
				labels_succ[7].SetText("1")
			} else {
				labels_fail[7].SetText("1")
			}
		})
	}
}

func ROUND(index int, bits []bool, buf []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	if index == 0 {
		//"单比特频数检测"
		p, q := randomness.MonoBitFrequencyTest(bits)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 1 {
		//"块内频数检测 m=10000" 10  100 1000 10000 1000000 根据长度自动选择
		p, q := randomness.FrequencyWithinBlockTest(bits)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 2 {
		//"扑克检测 m=4"  至少8字节
		//"扑克检测 m=8"
		p, q := randomness.PokerProto(bits, 4)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 3 {
		p, q := randomness.PokerProto(bits, 8)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 4 || index == 5 {
		//"重叠子序列检测 m=3 P1"
		//"重叠子序列检测 m=3 P2"
		//"重叠子序列检测 m=5 P1"
		//"重叠子序列检测 m=5 P2"
		p1, p2, q1, q2 := randomness.OverlappingTemplateMatchingProto(bits, 3)
		if index == 4 {
			Perr[index] = p1
			Qerr[index] = q1
		}
		if index == 5 {
			Perr[index] = p2
			Qerr[index] = q2
		}
	}

	if index == 6 || index == 7 {
		p1, p2, q1, q2 := randomness.OverlappingTemplateMatchingProto(bits, 5)
		if index == 6 {
			Perr[index] = p1
			Qerr[index] = q1
		}
		if index == 7 {
			Perr[index] = p2
			Qerr[index] = q2
		}
	}

	if index == 8 {
		//"游程总数检测"
		p, q := randomness.RunsTest(bits)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 9 {
		//"游程分布检测"
		p, q := randomness.RunsDistributionTest(bits)
		Perr[index] = p
		Qerr[index] = q
	}

	//至少128字节
	if index == 10 {
		//"块内最大1游程检测 m=10000"
		p, q := randomness.LongestRunOfOnesInABlockTest(bits, true)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 11 {
		//"块内最大0游程检测 m=10000"
		p, q := randomness.LongestRunOfOnesInABlockTest(bits, false)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 12 {
		//"二元推导检测 k=3" 至少7字节
		p, q := randomness.BinaryDerivativeProto(bits, 3)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 13 {
		//"二元推导检测 k=7"
		p, q := randomness.BinaryDerivativeProto(bits, 7)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 14 {
		//"自相关检测 d=1"
		//"自相关检测 d=2"
		//"自相关检测d=8"
		//"自相关检测 d=16" 至少16字节
		p, q := randomness.AutocorrelationProto(bits, 1)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 15 {
		p, q := randomness.AutocorrelationProto(bits, 2)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 16 {
		p, q := randomness.AutocorrelationProto(bits, 8)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 17 {
		p, q := randomness.AutocorrelationProto(bits, 16)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 18 {
		//"矩阵秩检测"
		p, q := randomness.MatrixRankTest(bits)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 19 {
		//"累加和检测"  前向
		p, q := randomness.CumulativeTest(bits, true)
		Perr[index] = p
		Qerr[index] = q
	}
	if index == 20 {
		//"累加和检测"  后向
		p, q := randomness.CumulativeTest(bits, false)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 21 {
		//"近似熵检测 m=2"
		//"近似熵检测 m=5"
		p, q := randomness.ApproximateEntropyProto(bits, 2)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 22 {
		p, q := randomness.ApproximateEntropyProto(bits, 5)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 23 {
		//"线性复杂度检测 m=500"
		//"线性复杂度检测 m=1000"
		p, q := randomness.LinearComplexityProto(bits, 500)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 24 {
		p, q := randomness.LinearComplexityProto(bits, 1000)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 25 {
		//"Maurer通用统计检测 L=7 Q=1280"
		p, q := randomness.MaurerUniversalTest(bits)
		Perr[index] = p
		Qerr[index] = q
	}

	if index == 26 {
		//"离散傅里叶检测"
		//p, q := randomness.DiscreteFourierTransformTest(bits)
		p, q := fun_DiscreteFourierTransformTest(buf, len(bits))
		Perr[index] = p
		Qerr[index] = q
	}

	//fmt.Printf("index=%d %s\n", index, itemStr[index])
	ui.QueueMain(func() {
		if Perr[index] >= 0.01 {
			labels_succ[index].SetText("1")
		} else {
			labels_fail[index].SetText("1")
		}
		labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
		prosessbar.SetValue(int(math.Min(float64(prosessbar.Value()+4), 97)))
	})
}

func worker(jobs <-chan string, out chan<- *R) {
	for filename := range jobs {
		buf, _ := os.ReadFile(filename)
		bits := randomness.B2bitArr(buf)
		//buf = nil

		arr := make([]float64, 0, len(selectslice))
		arrQ := make([]float64, 0, len(selectslice))

		if 1 == setcnt && len(bits) >= 167772160 {

			//单文件大于20MB
			var wg sync.WaitGroup

			tNum := 0
			for i := 0; i < len(selectslice); i++ {

				if 4 == selectslice[i] && i < len(selectslice)-1 && 5 == selectslice[i+1] {
					go ROUNDOverlapping(3, bits, &wg)
					i++
					wg.Add(1)
					tNum = tNum + 1
				} else if 6 == selectslice[i] && i < len(selectslice)-1 && 7 == selectslice[i+1] {
					go ROUNDOverlapping(5, bits, &wg)
					i++
					wg.Add(1)
					tNum = tNum + 1
				} else {
					go ROUND(selectslice[i], bits, buf, &wg)
					wg.Add(1)
					tNum = tNum + 1
				}

				if i > 0 && 0 == tNum%8 {
					wg.Wait()
					fmt.Printf("------fo wait----\n")
				}
			}

			wg.Wait()
			fmt.Printf("------wait----\n")

			buf = nil
			//fmt.Printf("Perr=%v\n", Perr)

			for i := 0; i < len(selectslice); i++ {
				arr = append(arr, Perr[selectslice[i]])
				arrQ = append(arrQ, Qerr[selectslice[i]])
			}

			go func(file string) {
				out <- &R{path.Base(file), arr, arrQ}
			}(filename)

		} else {

			if selectArry[0] == 1 {
				//"单比特频数检测"
				p, q := randomness.MonoBitFrequencyTest(bits)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}

			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(4)
				})
			}

			if selectArry[1] == 1 {
				//"块内频数检测 m=10000" 10  100 1000 10000 1000000 根据长度自动选择
				p, q := randomness.FrequencyWithinBlockTest(bits)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(8)
				})
			}

			if selectArry[2] == 1 {
				//"扑克检测 m=4"  至少8字节
				//"扑克检测 m=8"
				p, q := randomness.PokerProto(bits, 4)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(12)
				})
			}

			if selectArry[3] == 1 {
				p, q := randomness.PokerProto(bits, 8)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(16)
				})
			}

			if selectArry[4] == 1 || selectArry[5] == 1 {
				//"重叠子序列检测 m=3 P1"
				//"重叠子序列检测 m=3 P2"
				//"重叠子序列检测 m=5 P1"
				//"重叠子序列检测 m=5 P2"
				p1, p2, q1, q2 := randomness.OverlappingTemplateMatchingProto(bits, 3)
				if selectArry[4] == 1 {
					arr = append(arr, p1)
					arrQ = append(arrQ, q1)
				}
				if selectArry[5] == 1 {
					arr = append(arr, p2)
					arrQ = append(arrQ, q2)
				}
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(20)
				})
			}

			if selectArry[6] == 1 || selectArry[7] == 1 {
				p1, p2, q1, q2 := randomness.OverlappingTemplateMatchingProto(bits, 5)
				if selectArry[6] == 1 {
					arr = append(arr, p1)
					arrQ = append(arrQ, q1)
				}
				if selectArry[7] == 1 {
					arr = append(arr, p2)
					arrQ = append(arrQ, q2)
				}
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(24)
				})
			}

			if selectArry[8] == 1 {
				//"游程总数检测"
				p, q := randomness.RunsTest(bits)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(28)
				})
			}

			if selectArry[9] == 1 {
				//"游程分布检测"
				p, q := randomness.RunsDistributionTest(bits)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(32)
				})
			}

			//至少128字节
			if selectArry[10] == 1 {
				//"块内最大1游程检测 m=10000"
				p, q := randomness.LongestRunOfOnesInABlockTest(bits, true)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(36)
				})
			}

			if selectArry[11] == 1 {
				//"块内最大0游程检测 m=10000"
				p, q := randomness.LongestRunOfOnesInABlockTest(bits, false)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(40)
				})
			}

			if selectArry[12] == 1 {
				//"二元推导检测 k=3" 至少7字节
				p, q := randomness.BinaryDerivativeProto(bits, 3)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(44)
				})
			}

			if selectArry[13] == 1 {
				//"二元推导检测 k=7"
				p, q := randomness.BinaryDerivativeProto(bits, 7)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(48)
				})
			}

			if selectArry[14] == 1 {
				//"自相关检测 d=1"
				//"自相关检测 d=2"
				//"自相关检测d=8"
				//"自相关检测 d=16" 至少16字节
				p, q := randomness.AutocorrelationProto(bits, 1)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(52)
				})
			}

			if selectArry[15] == 1 {
				p, q := randomness.AutocorrelationProto(bits, 2)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(56)
				})
			}

			if selectArry[16] == 1 {
				p, q := randomness.AutocorrelationProto(bits, 8)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(60)
				})
			}

			if selectArry[17] == 1 {
				p, q := randomness.AutocorrelationProto(bits, 16)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(64)
				})
			}

			if selectArry[18] == 1 {
				//"矩阵秩检测"
				p, q := randomness.MatrixRankTest(bits)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(68)
				})
			}

			if selectArry[19] == 1 {
				//"累加和检测"  前向
				p, q := randomness.CumulativeTest(bits, true)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(72)
				})
			}

			if selectArry[20] == 1 {
				//"累加和检测"  后向
				p, q := randomness.CumulativeTest(bits, false)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(80)
				})
			}

			if selectArry[21] == 1 {
				//"近似熵检测 m=2"
				//"近似熵检测 m=5"
				p, q := randomness.ApproximateEntropyProto(bits, 2)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(84)
				})
			}

			if selectArry[22] == 1 {
				p, q := randomness.ApproximateEntropyProto(bits, 5)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(88)
				})
			}

			if selectArry[23] == 1 {
				//"线性复杂度检测 m=500"
				//"线性复杂度检测 m=1000"
				p, q := randomness.LinearComplexityProto(bits, 500)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(92)
				})
			}

			if selectArry[24] == 1 {
				p, q := randomness.LinearComplexityProto(bits, 1000)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(96)
				})
			}

			if selectArry[25] == 1 {
				//"Maurer通用统计检测 L=7 Q=1280"
				p, q := randomness.MaurerUniversalTest(bits)
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}
			if 1 == setcnt {
				ui.QueueMain(func() {
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
					prosessbar.SetValue(98)
				})
			}

			if selectArry[26] == 1 {
				//"离散傅里叶检测"
				//p, q := randomness.DiscreteFourierTransformTest(bits)
				p, q := fun_DiscreteFourierTransformTest(buf, len(bits))
				buf = nil
				arr = append(arr, p)
				arrQ = append(arrQ, q)
			}

			go func(file string) {
				out <- &R{path.Base(file), arr, arrQ}
			}(filename)
		}
	}
}

var distributions [][]float64

func createDistributions(s, m int) [][]float64 {
	res := make([][]float64, m)
	for i := 0; i < m; i++ {
		res[i] = make([]float64, 0, s)
	}
	return res
}

// 结果集写入文件工作器
func resultWriter(in <-chan *R, w io.StringWriter, cnt []int32, wg *sync.WaitGroup) {
	for r := range in {
		_, _ = w.WriteString(r.Name)

		// ThresholdQ 样本分布均匀性 (k=10)
		//Pt := detect.ThresholdQ(r.Q)

		for j := 0; j < len(r.P); j++ {
			if r.P[j] >= 0.01 {
				atomic.AddInt32(&cnt[j], 1)
			} else {
				atomic.AddInt32(&failcnt[j], 1)
			}
			_, _ = w.WriteString(fmt.Sprintf(", %0.6f|%0.6f", r.P[j], r.Q[j]))

			distributions[j] = append(distributions[j], r.Q[j])
		}
		_, _ = w.WriteString("\n")
		//_, _ = w.WriteString(fmt.Sprintf(", %0.6f\n", Pt))

		wg.Done()
	}
}

func StrToGBK(str string) string {
	sysType := runtime.GOOS
	if sysType == "windows" {
		//将utf-8编码的字符串转换为GBK编码
		ret, _ := simplifiedchinese.GBK.NewEncoder().String(str)
		return ret //如果转换失败返回空字符串
	} else {
		return str
	}

	//如果是[]byte格式的字符串，可以使用Bytes方法
	//b, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
	//return string(b)
}

var startTime time.Time
var processfile int32
var percent int32

func makeBasicControlsPage() ui.Control {

	Alpha = 0.01

	//垂直布局
	hboxMain := ui.NewHorizontalBox()
	hboxMain.SetPadded(true)

	//水平布局
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hboxMain.Append(vbox, true)

	//垂直布局
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, true)

	//水平布局左
	vboxleft := ui.NewVerticalBox()
	vboxleft.SetPadded(false)
	hbox.Append(vboxleft, false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	//水平布局右
	vboxright := ui.NewVerticalBox()
	vboxright.SetPadded(false)
	hbox.Append(vboxright, true)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vboxstat := ui.NewVerticalBox()
	vboxstat.SetPadded(false)
	hbox.Append(vboxstat, true)

	vboxleft.Append(ui.NewLabel("检测项"), false)
	vboxright.Append(ui.NewLabel(">=显著水平样本数"), false)
	vboxstat.Append(ui.NewLabel("<显著水平样本数"), false)

	seletctAll = true
	for i := 0; i < len(itemStr); i++ {
		checkboxs[i] = ui.NewCheckbox(itemStr[i])
		checkboxs[i].SetChecked(seletctAll)
		labels_succ[i] = ui.NewLabel("0")
		labels_fail[i] = ui.NewLabel("0")

		vboxleft.Append(checkboxs[i], true)
		vboxleft.Append(ui.NewHorizontalSeparator(), false)
		vboxright.Append(labels_succ[i], true)
		vboxright.Append(ui.NewHorizontalSeparator(), false)
		vboxstat.Append(labels_fail[i], true)
		vboxstat.Append(ui.NewHorizontalSeparator(), false)
	}

	buttonSelectAll := ui.NewButton("")
	if true == seletctAll {
		buttonSelectAll.SetText("反选")
	} else if false == seletctAll {
		buttonSelectAll.SetText("全选")
	}

	buttonSelectAll.OnClicked(func(*ui.Button) {

		seletctAll = !seletctAll
		for i := 0; i < len(itemStr); i++ {
			checkboxs[i].SetChecked(seletctAll)
		}
		if true == seletctAll {
			buttonSelectAll.SetText("全选")
		} else if false == seletctAll {
			buttonSelectAll.SetText("反选")
		}
	})

	hboxAl := ui.NewHorizontalBox()
	hboxAl.SetPadded(true)
	vbox.Append(hboxAl, false)

	hboxAl.Append(buttonSelectAll, true)
	hboxAl.Append(ui.NewLabel("显著水平:"), false)
	entryAlpha = ui.NewEntry()
	entryAlpha.SetReadOnly(true)
	entryAlpha.SetText(fmt.Sprintf("%1.3f", Alpha))
	hboxAl.Append(entryAlpha, false)

	hboxAl.Append(ui.NewLabel("分布均匀性的显著性水平:"), false)
	entryAlphaT := ui.NewEntry()
	entryAlphaT.SetReadOnly(true)
	entryAlphaT.SetText(fmt.Sprintf("%1.4f", randomness.AlphaT))
	hboxAl.Append(entryAlphaT, false)

	//右侧视图
	vboxOp := ui.NewVerticalBox()
	vboxOp.SetPadded(true)
	hboxMain.Append(ui.NewVerticalSeparator(), false)
	hboxMain.Append(vboxOp, true)

	//进度条
	prosessbar = ui.NewProgressBar()
	prosessbar.SetValue(0)
	//vbox.Append(prosessbar, false)
	vboxOp.Append(prosessbar, true)

	//其他
	vboxOp.Append(ui.NewLabel("建议使用1000*1000 bit 或者128kb大小的样本进行检测。"), false)
	hbox1 := ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("样本大小:"), true)
	labelsetbit = ui.NewLabel("")
	hbox1.Append(labelsetbit, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("样本数量:"), true)
	labelsetNum = ui.NewLabel("")
	hbox1.Append(labelsetNum, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("成功通过检测项,需要>=显著水平的样本数量:"), true)
	labelsetAlphaNum = ui.NewLabel("")
	hbox1.Append(labelsetAlphaNum, true)
	vboxOp.Append(hbox1, false)

	//添加分割线
	//vbox.Append(ui.NewHorizontalSeparator(), false)
	vboxOp.Append(ui.NewHorizontalSeparator(), false)

	//样本路径
	hboxpath := ui.NewHorizontalBox()
	hboxpath.SetPadded(true)

	entryOutPath := ui.NewEntry()

	buttonFile = ui.NewButton("选择样本文件")
	entryPath = ui.NewEntry()
	entryPath.SetReadOnly(true)
	buttonFile.OnClicked(func(*ui.Button) {
		setcnt = 0
		entryPath.SetText("")
		entryOutPath.SetText("")
		filename := ui.OpenFile(mainwin)
		if filename != "" {
			files, err := ioutil.ReadDir(filepath.Dir(filename))
			if err == nil {
				for _, file := range files {

					if strings.HasSuffix(file.Name(), ".bin") || strings.HasSuffix(file.Name(), ".dat") {
						setcnt++
					}
				}
			}
		}

		if setcnt == 0 {
			ui.MsgBoxError(mainwin, "随机数检测工具", "样本文件必须是 .dat 或 .bin文件")
			labelsetAlphaNum.SetText("")
			labelsetNum.SetText("")
		} else {

			buf, _ := os.ReadFile(filename)
			labelsetbit.SetText(fmt.Sprintf("%dbit %dbyte", len(buf)*8, len(buf)))

			if len(buf) < (100 * 1024) {
				//if len(buf) < (7 * 1280 / 8) {
				//checkboxs[25].SetChecked(false)
				//ui.MsgBoxError(mainwin, "随机数检测工具", fmt.Sprintf("Maurer通用统计检测 数据长度至少要满足 L*Q %dbyte", 7*1280/8))
				ui.MsgBoxError(mainwin, "随机数检测工具", "请提供高质量随机数")
				setcnt = 0
				//checkboxs[len(itemStr)-1].SetChecked(true)
			}
			/*else if len(buf) > 1048576 {
				ui.MsgBoxError(mainwin, "随机数检测工具", "离散傅里叶检测 最大支持128k 文件,")
				checkboxs[len(itemStr)-1].SetChecked(false)
				//setcnt = 0
			} else {
				//checkboxs[len(itemStr)-1].SetChecked(true)
			}*/

			if setcnt > 0 {
				if setcnt == 1 {
					entryPath.SetText(filename)
				} else {
					entryPath.SetText(filepath.Dir(filename))
				}
				outpath := path.Join(filepath.Dir(filepath.Dir(filename)), "/LMRandomCheckReport"+time.Now().Format("-20060102-15-04-05")+".csv")
				//outpath := filepath.Dir(filename) + "/RandomnessTestReport.csv"
				_ = os.MkdirAll(filepath.Dir(outpath), os.FileMode(0600))

				entryOutPath.SetText(outpath)
				//var num float64 = (1 - Alpha - 3*math.Sqrt(float64((Alpha*(1-Alpha))/float64(setcnt)))) * float64(setcnt)
				//fmt.Printf("%f %d", num, int(math.Ceil(num)))
				//labelsetAlphaNum.SetText(fmt.Sprintf("%d", int(math.Ceil(num))))

				labelsetAlphaNum.SetText(fmt.Sprintf("%d", detect.Threshold(setcnt)))
				labelsetNum.SetText(fmt.Sprintf("%d", setcnt))
			}
		}
	})
	hboxpath.Append(ui.NewLabel("样本路径:"), false)
	hboxpath.Append(entryPath, true)
	hboxpath.Append(buttonFile, false)

	//vbox.Append(hboxpath, false)
	vboxOp.Append(hboxpath, false)

	//报告路径
	hboxpath = ui.NewHorizontalBox()
	hboxpath.SetPadded(true)

	buttonOK = ui.NewButton("开 始 检 测 ")
	buttonOK.OnClicked(func(*ui.Button) {

		filename := entryOutPath.Text()

		if filename != "" {

			for j := 0; j < len(itemStr); j++ {
				labels_succ[j].SetText("0")
				labels_fail[j].SetText("0")
			}
			prosessbar.SetValue(0)
			labeltime.SetText("")
			labeltimeend.SetText("")
			labeltimebegin.SetText("")

			_ = os.MkdirAll(filepath.Dir(filename), os.FileMode(0600))

			selectslice = selectslice[0:0]
			for i := 0; i < len(itemStr); i++ {
				if checkboxs[i].Checked() {
					selectslice = append(selectslice, i)
					selectArry[i] = 1
				} else {
					selectArry[i] = 0
				}
			}
			for i := 0; i < len(failcnt); i++ {
				failcnt[i] = 0
			}

			inputPath = entryPath.Text()

			if setcnt == 1 {
				inputPath = filepath.Dir(inputPath)
			}
			reportPath = entryOutPath.Text()

			startTime = time.Now()
			labeltimebegin.SetText(startTime.Format("2006.01.02 15:04:05"))

			buttonOK.Disable()
			buttonFile.Disable()

			//线程
			go func() {

				n := runtime.NumCPU()
				out := make(chan *R)
				jobs := make(chan string)

				/*
					//8核心一下的cpu使用一半的线程数
					if n <= 8 {
						n = n / 2
					}
				*/
				w, err := os.OpenFile(reportPath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(0600))
				if err != nil {
					fmt.Printf(">> 无法打开写入报告文件")
					ui.QueueMain(func() {
						ui.MsgBoxError(mainwin, "随机数检测工具", "无法打开写入文件 "+reportPath)
						buttonOK.Enable()
						buttonFile.Enable()
					})
					return
				}
				defer w.Close()

				var reportheader string
				reportheader = "源数据"

				for i := 0; i < len(selectslice); i++ {
					reportheader += ","
					reportheader += itemStr[selectslice[i]]
				}
				reportheader += "\n"

				reportheader = StrToGBK(reportheader)
				_, _ = w.WriteString(reportheader)

				files, err := ioutil.ReadDir(inputPath)
				if err != nil {
					ui.QueueMain(func() {
						ui.MsgBoxError(mainwin, "随机数检测工具", "读取随机数目录失败"+inputPath)
						buttonOK.Enable()
						buttonFile.Enable()
					})

					return
				}

				distributions = createDistributions(setcnt, len(selectslice))

				var wg sync.WaitGroup
				cnt = make([]int32, len(selectslice))

				wg.Add(setcnt)

				// 启动数据写入消费者
				go resultWriter(out, w, cnt, &wg)
				// 检测工作器
				for i := 0; i < n; i++ {
					go worker(jobs, out)
				}

				processfile = 0
				if setcnt > 100 {
					percent = int32(setcnt / 100)
				}

				// 结果工作器
				for _, file := range files {
					if strings.HasSuffix(file.Name(), ".bin") || strings.HasSuffix(file.Name(), ".dat") {

						atomic.AddInt32(&processfile, 1)

						if setcnt > 1 {
							ui.QueueMain(func() {
								//减少刷新次数
								if 0 == processfile%10 {
									prosessbar.SetValue(int(math.Min(float64(processfile/percent), 97)))
									labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
									for j := 0; j < len(selectslice); j++ {
										labels_succ[selectslice[j]].SetText(fmt.Sprintf("%d", cnt[j]))
										labels_fail[selectslice[j]].SetText(fmt.Sprintf("%d", failcnt[j]))
									}
								}

							})

						} else {
							ui.QueueMain(func() {
								prosessbar.SetValue(2)
							})
						}

						jobs <- (path.Join(inputPath, "/", file.Name()))
					}
				}

				wg.Wait()

				ui.QueueMain(func() {

					elapsedTime := time.Since(startTime) / time.Second // duration in s
					labeltime.SetText(fmt.Sprintf("%d 分钟 %d 秒", elapsedTime/60, elapsedTime%60))
					labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))

					for j := 0; j < len(selectslice); j++ {
						labels_succ[selectslice[j]].SetText(fmt.Sprintf("%d", cnt[j]))
						labels_fail[selectslice[j]].SetText(fmt.Sprintf("%d", failcnt[j]))
					}

					var errStr string
					var numCount int = 0
					//通过的组数
					//var numpass int32 = int32(math.Ceil((1 - Alpha - 3*math.Sqrt(float64((Alpha*(1-Alpha))/float64(setcnt)))) * float64(setcnt)))
					var numpass int32 = int32(detect.Threshold(setcnt))
					for j := 0; j < len(selectslice); j++ {

						if cnt[j] < numpass {
							numCount++
							errStr += itemStr[selectslice[j]]
							errStr += ",\n"
						}
					}

					prosessbar.SetValue(100)
					buttonOK.Enable()
					buttonFile.Enable()

					errStr += "\n"
					for i := 0; i < len(distributions); i++ {
						Pt := detect.ThresholdQ(distributions[i])

						if Pt < randomness.AlphaT {
							numCount++
							errStr += itemStr[selectslice[i]]
							errStr += "均匀性较差,\n"
						}
					}

					if 0 == numCount {
						ui.MsgBox(mainwin, "成功", "所有检测项全部通过")
					} else {
						ui.MsgBoxError(mainwin, "失败", fmt.Sprintf("有%d项检测未通过:\n%s\n", numCount, errStr))
					}

					processfile = 0
					percent = 1

				})

				reportheader = "分布均匀性"
				reportheader = StrToGBK(reportheader)
				_, _ = w.WriteString(reportheader)
				for i := 0; i < len(distributions); i++ {
					Pt := detect.ThresholdQ(distributions[i])

					_, _ = w.WriteString(fmt.Sprintf(", %0.6f", Pt))
				}
				_, _ = w.WriteString("\n")

				reportheader = "总计"
				reportheader = StrToGBK(reportheader)
				_, _ = w.WriteString(reportheader)
				for i := 0; i < len(cnt); i++ {
					_, _ = w.WriteString(fmt.Sprintf(", %d", cnt[i]))
				}
				_, _ = w.WriteString("\n")

			}()
		} else {
			ui.MsgBoxError(mainwin, "随机数检测工具", "请选择样本文件 ")
			return
		}
	})

	hboxpath.Append(ui.NewLabel("报告路径:"), false)
	hboxpath.Append(entryOutPath, true)
	hboxpath.Append(buttonOK, false)

	//vbox.Append(hboxpath, false)
	vboxOp.Append(hboxpath, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("开始测试时间:"), true)
	labeltimebegin = ui.NewLabel("")
	hbox1.Append(labeltimebegin, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("测试完成时间:"), true)
	labeltimeend = ui.NewLabel("")
	hbox1.Append(labeltimeend, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("耗时:"), true)
	labeltime = ui.NewLabel("")
	hbox1.Append(labeltime, true)
	vboxOp.Append(hbox1, false)

	grid := ui.NewGrid()
	grid.SetPadded(false)

	grid.Append(ui.NewLabel("国密随机数质量检测工具,支持《GMT 0005-2021 随机性检测规范》。"),
		0, 0, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)
	grid.Append(ui.NewLabel("利用多核CPU进行检测,大大缩短检测时间。"),
		0, 1, 1, 1,
		true, ui.AlignCenter, false, ui.AlignCenter)

	grid.Append(ui.NewLabel("北京世纪龙脉科技有限公司 V1.7"),
		0, 2, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)

	vboxOp.Append(grid, true)

	return hboxMain
}

func setupUI() {
	mainwin = ui.NewWindow("随机数检测工具-北京世纪龙脉科技有限公司", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("GMT 0005-2021 随机性检测", makeBasicControlsPage())
	tab.SetMargined(0, true)

	mainwin.Show()

}

func main() {

	percent = 1
	ui.Main(setupUI)
}

//go 大于1.8
//Debian、Ubuntu 等：sudo apt-get install libgtk-3-dev
//Red Hat / Fedora 等：sudo dnf install gtk3-devel
//go get github.com/andlabs/libui
//go get github.com/andlabs/ui
//go get golang.org/x/text
//go get github.com/Trisia/randomness
//go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

//apt-get install golang-go
//wget https://dl.google.com/go/go1.19.3.linux-amd64.tar.gz
//tar -zxvf go1.19.3.linux-amd64.tar.gz
//解压到/home/lm/Desktop/
//export GOROOT=/home/lm/Desktop/go
//export GOPATH=/home/lm/go
//export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
//go version
//go mod init RandomCheckTool

//隐藏命令行
//go build -ldflags="-H windowsgui"  RandomCheckTool.go

//图标
//go get github.com/akavel/rsrc
//main.rc
//IDI_ICON1 ICON "logo.ico"
//windres -o main.syso main.rc
//windres -i rc/main.rc -O coff -o main.syso
//go build -o RandomCheckTool.exe  -ldflags="-H windowsgui"

//其它
//go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

//编译andlabs/libui
//apt-get install ninja-build
//pip3 install --user meson

//export PATH=$PATH:/root/.local/bin

//cd /usr/local
//wget --no-check-certificate https://www.python.org/ftp/python/3.6.5/Python-3.6.5.tgz
//wget --no-check-certificate https://www.python.org/ftp/python/3.7.5/Python-3.7.5.tgz
//tar -xzvf Python-3.6.5.tgz
//cd Python-3.6.5
//./configure --prefix=/usr/local/python3.6.5
//make
//make install
//cp /usr/bin/python3 /usr/bin/python3_bak
//rm /usr/bin/python3
//ln -s /usr/local/Python-3.6.5/python /usr/bin/python3

//wget https://bootstrap.pypa.io/pip/3.6/get-pip.py
//python3 get-pip.py

//grep 取行，awk 按条件取指定列，cut 按分隔符取指定列。
//sed主要是用来将数据进行选取、替换、删除、新增的命令。可以放在管道符之后处理。

// ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1
//cp $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 ) lib/
//ll $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 )
//ll $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 ) | awk '{print $11,$10,$9}' | cut -d '>' -f 2
